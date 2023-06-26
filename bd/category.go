package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	//"strconv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/livingstone23/gambit/models"
	"github.com/livingstone23/gambit/tools"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Inicializando funcion  db.InsertCategory")

	err := DBConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentencia := "INSERT INTO category () Values('" + c.CategName + "','" + c.CategPath + "')"

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

	fmt.Println("Insert Category > Ejecucion Exitosa")
	return LastInsertId, nil

}

func UpdateCategory(c models.Category) error {
	fmt.Println("Inicializando funcion  db.UpdateCategory")

	err := DBConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "Update category SET "
	if len(c.CategName) > 0 {
		sentencia += " Categ_Name = '" + tools.EscapeString(c.CategName) + "'"
	}

	if len(c.CategPath) > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia += " , "
		}
		sentencia += " Categ_Path = '" + tools.EscapeString(c.CategPath) + "'"
	}

	sentencia += " Where Categ_Id = " + strconv.Itoa(c.CategID)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Category > ejecutado Exitosamente")
	return nil
}
