package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/livingstone23/gambit/models"
)

func InsertOrder(o models.Orders) (int64, error) {
	fmt.Println("Inicializando funcion  bd.InsertOrder")

	err := DBConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentencia := "INSERT INTO orders (Order_UserUUID, Order_Total, Order_AddId) Values ('"
	sentencia += o.Order_UserUUID + "'," + strconv.FormatFloat(o.Order_Total, 'f', -1, 64) + "," + strconv.Itoa(o.Order_AddId) + ")"

	fmt.Println("Vamos a ejecutar la sentancia: ")
	fmt.Println(sentencia)

	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err
	}

	for _, od := range o.OrdersDetails {
		sentencia = "INSERT INTO orders_detail (OD_OrderId, OD_ProdId,  OD_Quantity, OD_Price) VALUES (" + strconv.Itoa(int(LastInsertId))
		sentencia += "," + strconv.Itoa(od.OD_ProdId) + "," + strconv.Itoa(od.OD_Quantity) + "," + strconv.FormatFloat(float64(od.OD_Price), 'f', -1, 64) + ")"
		fmt.Println(sentencia)

		_, err = Db.Exec(sentencia)
		if err != nil {
			fmt.Println(err.Error())
			return 0, err
		}
	}

	fmt.Println("Insert order > Ejecucion exitosa ")
	return LastInsertId, nil

}

func SelectOrders(user string, fechaDesde string, fechaHasta string, page int, orderId int) ([]models.Orders, error) {
	fmt.Println("Inicializando funcion  bd.SelectOrders")

	var Orders []models.Orders

	sentencia := "SELECT Order_Id, Order_UserUUID, Order_AddId, Order_Date, Order_Total FROM orders "

	if orderId > 0 {
		sentencia += " WHERE Order_Id = " + strconv.Itoa(orderId)
	} else {
		offset := 0
		if page == 0 {
			page = 1
		}
		if page > 1 {
			offset = (10 * (page - 1))
		}

		if len(fechaHasta) == 10 {
			fechaHasta += " 23:59:59"
		}

		var where string
		var whereUser string = " Order_UserUUID = '" + user + "'"

		if len(fechaDesde) > 0 && len(fechaHasta) > 0 {
			where += " WHERE Order_Date BETWEEN '" + fechaDesde + "' AND '" + fechaHasta
		}

		if len(where) > 0 {
			where += " AND " + whereUser
		} else {
			where += " WHERE " + whereUser
		}

		limit := " LIMIT 10 "
		if offset > 0 {
			limit += " OFFSET " + strconv.Itoa(offset)
		}

		sentencia += where + limit
	}

	fmt.Println("Vamos a ejecutar la sentencia")
	fmt.Println(sentencia)

	err := DBConnect()
	if err != nil {
		return Orders, err
	}

	defer Db.Close()

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return Orders, err
	}

	//Nos aseguramos que cierra la sentencia
	defer rows.Close()

	for rows.Next() {
		var Order models.Orders
		//var OrderDate sql.NullTime
		var OrderAddId sql.NullInt32
		err := rows.Scan(&Order.Order_Id, &Order.Order_UserUUID, &OrderAddId, &Order.Order_Date, &Order.Order_Total)
		if err != nil {
			return Orders, err
		}
		//Order.Order_Date = OrderDate.Time.String()
		Order.Order_AddId = int(OrderAddId.Int32)

		var rowsD *sql.Rows
		sentenciaD := "SELECT OD_Id, Od_ProdId, OD_Quantity, OD_Price FROM orders_detail WHERE OD_OrderID =  " + strconv.Itoa(Order.Order_Id)

		fmt.Println("Vamos a ejecutar la sentenciaD")
		fmt.Println(sentenciaD)

		rowsD, err = Db.Query(sentenciaD)
		if err != nil {
			fmt.Println(err.Error())
			return Orders, err
		}

		for rowsD.Next() {
			var OD_Id int64
			var OD_ProdId int64
			var OD_Quantity int64
			var OD_Price float64

			err = rowsD.Scan(&OD_Id, &OD_ProdId, &OD_Quantity, &OD_Price)
			if err != nil {
				return Orders, err
			}

			var od models.OrdersDetails

			od.OD_Id = int(OD_Id)
			od.OD_ProdId = int(OD_ProdId)
			od.OD_Quantity = int(OD_Quantity)
			od.OD_Price = OD_Price

			Order.OrdersDetails = append(Order.OrdersDetails, od)
		}
		Orders = append(Orders, Order)
		rowsD.Close()
	}

	fmt.Println("Select Orders > Ejecucion exitosa")
	return Orders, nil
}
