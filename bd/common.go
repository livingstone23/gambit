package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/livingstone23/gambit/models"
	"github.com/livingstone23/gambit/secretm"
)

var SecretModel models.SecretRDSJson
var err error

var Db *sql.DB

func ReadSecret() error {

	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DBConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexion exitosa de la BD")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = claves.Dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(userUUIID string) (bool, string) {
	fmt.Println("Inicializando funcion  UserIsAdmin")

	err := DBConnect()
	if err != nil {
		return false, err.Error()
	}
	//Cierra conexion a la base de datos
	defer Db.Close()

	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + userUUIID + "' AND User_Status = 0"
	fmt.Println(sentencia)

	rows, err := Db.Query(sentencia)
	if err != nil {
		return false, err.Error()
	}

	var valor string
	//Nos posicionamos en el primer registro, para poder leer los datos
	rows.Next()
	//con esta sintaxis si obtengo valor de consulta lo guardo en variable
	rows.Scan(&valor)

	fmt.Println("UserIdAdmin > Ejecucion exitosa - valor devuelto " + valor)
	if valor == "1" {
		return true, ""
	}

	return false, "User is not admin"

}

func UserExists(UserUUID string) (error, bool) {
	fmt.Println("Inicializando funcion  db.UserExists")

	err := DBConnect()
	if err != nil {
		return err, false
	}

	defer Db.Close()

	sentencia := "SELECT 1 FROM users WHERE User_UUID='" + UserUUID + "'"
	fmt.Println(sentencia)

	rows, err := Db.Query(sentencia)
	if err != nil {
		return err, false
	}

	var valor string
	rows.Next()
	//con esta sintaxis si obtengo valor de consulta lo guardo en variable
	rows.Scan(&valor)

	fmt.Println("UserExist Ejecucion exitosa - Valor devuelto " + valor)

	if valor == "1" {
		return nil, true
	}
	return nil, false

}
