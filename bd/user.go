package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/livingstone23/gambit/models"
	"github.com/livingstone23/gambit/tools"
)

func UpdateUser(UField models.User, User string) error {
	fmt.Println("Inicializando funcion  db.UpdateUser")

	err := DBConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "UPDATE users SET "

	coma := ""
	if len(UField.UserFirstName) > 0 {
		sentencia = ","
		sentencia += coma + "User_FirstName = '" + UField.UserFirstName + "'"
	}

	if len(UField.UserLastName) > 0 {
		sentencia += coma + "User_LastName = '" + UField.UserLastName + "'"
	}

	sentencia += ", User_DateUpg = '" + tools.FechaMySQL() + "' Where User_UUID='" + User + "'"

	println("ejecutando sentencia")
	println(sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update de User > Ejecucion exitosa")
	return nil

}

func SelectUser(UserId string) (models.User, error) {
	fmt.Println("Inicializando funcion  db.SelectUser")

	User := models.User{}

	err := DBConnect()
	if err != nil {
		return User, err
	}

	defer Db.Close()

	sentencia := "Select * from users Where User_UUID = '" + UserId + "'"

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	//Nos aseguramos que cierra la sentencia
	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return User, err
	}

	rows.Next()

	var firstName sql.NullString
	var lastName sql.NullString
	var dateUpg sql.NullTime

	err = rows.Scan(&User.UserUUID, &User.UserEmail, &firstName, &lastName, &User.UserStatus, &User.UserDateAdd, &dateUpg)

	User.UserFirstName = firstName.String
	User.UserLastName = lastName.String
	User.UserDateUpd = dateUpg.Time.String()

	fmt.Println("Select User > Ejecucion exitosa")
	return User, nil

}

func SelectUsers(Page int) (models.ListUsers, error) {
	fmt.Println("Inicializando funcion  db.SelectUsers")

	var listUser models.ListUsers
	User := []models.User{}

	err := DBConnect()
	if err != nil {
		return listUser, err
	}

	defer Db.Close()

	var offset int = (Page * 10) - 10
	var sentencia string
	var sentenciaCount string = "Select count(*) as registros from users"

	if offset > 0 {
		sentencia = "select * from users limit 10 OFFSET " + strconv.Itoa(offset)
	}

	if offset > 0 {
		sentencia += " OFFSET " + strconv.Itoa(offset)
	}

	var rowsCount *sql.Rows

	fmt.Print("Ejecutando sentenciaCount")
	fmt.Println(sentenciaCount)
	rowsCount, err = Db.Query(sentenciaCount)
	if err != nil {
		return listUser, err
	}

	defer rowsCount.Close()
	rowsCount.Next()

	var registros int
	rowsCount.Scan(&registros)
	listUser.TotalItems = registros

	var rows *sql.Rows
	rows, err = Db.Query(sentencia)
	if err != nil {
		return listUser, err
	}

	for rows.Next() {
		var u models.User

		var firstName sql.NullString
		var lastName sql.NullString
		var dateUpg sql.NullTime

		err = rows.Scan(&u.UserUUID, &u.UserEmail, &firstName, &lastName, &u.UserStatus, &u.UserDateAdd, &dateUpg)

		u.UserFirstName = firstName.String
		u.UserLastName = lastName.String
		u.UserDateUpd = dateUpg.Time.String()
		User = append(User, u)
	}

	fmt.Println("Select Users > Ejecucion exitosa")
	listUser.Data = User
	return listUser, nil

}
