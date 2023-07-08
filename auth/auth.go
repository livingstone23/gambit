package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

//Estructura que viaja en el token
type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidoToken(token string) (bool, error, string) {
	fmt.Println("Ingresando a la funcion ValidoToken")
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		fmt.Println("El token no es valido")
		return false, nil, "El token no es valido"
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("No se puede decodificar la parte del token : ", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("No se puede decodificar en la estructura JSON", err.Error())
		return false, err, err.Error()
	}

	ahora := time.Now()
	//creo una variable para poder comparar las fechas
	tm := time.Unix(int64(tkj.Exp), 0)

	//Funcion para comparar fechas
	if tm.Before(ahora) {
		fmt.Println("Fecha de expiracion token" + tm.String())
		fmt.Println("Token expirado !")
		return false, err, "Token expierado !!"
	}

	return true, nil, string(tkj.Username)

}
