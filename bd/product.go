package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/livingstone23/gambit/models"
	"github.com/livingstone23/gambit/tools"
)

func InsertProduct(p models.Product) (int64, error) {
	fmt.Println("Inicializando funcion  db.InsertProduct")

	err := DBConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentencia := "INSERT INTO product (Prod_Title"

	if len(p.ProdDescription) > 0 {
		sentencia += ", Prod_Description"
	}

	if p.ProdPrice > 0 {
		sentencia += ", Prod_Price"
	}

	if p.ProdCategId > 0 {
		sentencia += ", Prod_CategoryId"
	}

	if p.ProdStock > 0 {
		sentencia += ", Prod_Stock"
	}

	if len(p.ProdPath) > 0 {
		sentencia += ", Prod_Path"
	}

	sentencia += ") Values ('" + tools.EscapeString(p.ProdTitle) + "'"

	if len(p.ProdDescription) > 0 {
		sentencia += ", '" + tools.EscapeString(p.ProdDescription) + "'"
	}

	if p.ProdPrice > 0 {
		sentencia += ", '" + strconv.FormatFloat(p.ProdPrice, 'e', -1, 64)
	}

	if p.ProdCategId > 0 {
		sentencia += ", '" + strconv.Itoa(p.ProdStock)
	}

	if p.ProdStock > 0 {
		sentencia += ", Prod_Stock"
	}

	if len(p.ProdPath) > 0 {
		sentencia += ", '" + tools.EscapeString(p.ProdPath) + "'"
	}

	sentencia += ")"

	var result sql.Result

	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		fmt.Println(err.Error())
		return 0, err2
	}

	fmt.Println("Insert Product > Ejecucion Exitosa")
	return LastInsertId, nil
}

func UpdateProduct(p models.Product) error {
	fmt.Println("Inicializando funcion  db.UpdateProduct")

	err := DBConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "Update produc SET "

	sentencia = tools.ArmoSentencia(sentencia, "Prod_Title", "S", 0, 0, p.ProdTitle)
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Description", "S", 0, 0, p.ProdDescription)
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Price", "F", 0, p.ProdPrice, "")
	sentencia = tools.ArmoSentencia(sentencia, "Prod_CategoryId", "N", p.ProdCategId, 0, "")
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Stock", "N", p.ProdStock, 0, "")
	sentencia = tools.ArmoSentencia(sentencia, "Prod_Path", "S", 0, 0, p.ProdPath)

	sentencia += " Where Prod_Id = " + strconv.Itoa(p.ProdId)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Product > ejecutado Exitosamente")
	return nil
}

func DeleteProduct(id int) error {
	fmt.Println("Inicializando funcion  db.DeleteProduct")

	err := DBConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "Delete from product Where Prod_Id " + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Product > ejecutado Exitosamente")
	return nil
}
