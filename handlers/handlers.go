package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/livingstone23/gambit/auth"
	"github.com/livingstone23/gambit/routers"
)

func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Voy a procesar"+path, " > "+method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := validoAuthorization(path, method, headers)
	if !isOk {
		return statusCode, user
	}

	switch path[0:4] {
	case "user":
		return ProcesoUsers(body, path, method, user, id, request)

	case "prod":
		return ProcesoProducts(body, path, method, user, idn, request)

	case "stoc":
		return ProcesoStocks(body, path, method, user, idn, request)

	case "addr":
		return ProcesoAddress(body, path, method, user, idn, request)

	case "cate":
		return ProcesoCategory(body, path, method, user, idn, request)

	case "orde":
		return ProcesoOrders(body, path, method, user, idn, request)

	}

	return 400, "Method Invalid"

}

func validoAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	fmt.Println("Ingreando a metodo validoAuthorization")
	if (path == "product" && method == "GET") ||
		(path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token Requerido"
	}

	todoOK, err, msg := auth.ValidoToken(token)
	if !todoOK {
		if err != nil {
			fmt.Println("Error en el token" + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error en el token" + msg)
			return false, 401, msg
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg

}

func ProcesoUsers(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Ingreando a la funcion ProcesoUsers")

	return 400, "method invalid"
}

func ProcesoProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Ingreando a la funcion ProcesoProducts")

	return 400, "method invalid"
}

func ProcesoCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Ingreando a la funcion ProcesoCategory")

	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	case "PUT":
		return routers.UpdateCategory(body, user, id)
	}

	return 400, "method invalid"
}

func ProcesoStocks(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Ingreando a la funcion ProcesoStocks")

	return 400, "method invalid"
}

func ProcesoAddress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Ingreando a la funcion ProcesoAddress")

	return 400, "method invalid"
}

func ProcesoOrders(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Ingreando a la funcion ProcesoOrder")

	return 400, "method invalid"
}
