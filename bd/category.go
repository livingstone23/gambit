package bd

import (
	"database/sql"
	"fmt"

	//"strconv"
	//"strings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/livingstone23/gambit/models"
	//"github.com/livingstone23/gambit/tool"
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
