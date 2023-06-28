package routers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/livingstone23/gambit/bd"
	"github.com/livingstone23/gambit/models"
)

func InsertProduct(body string, User string) (int, string) {
	fmt.Println("Inicializando funcion  router.InsertProduct")

	var t models.Product

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especificar el Nombre (title) del producto "
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertProduct(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar registrar el producto" + t.ProdTitle + " > " + err2.Error()
	}

	return 200, "{ProdID: " + strconv.Itoa(int(result)) + "}"

}

func UpdateProduct(body string, User string, id int) (int, string) {
	fmt.Println("Inicializando funcion  router.UpdateProduct")

	var t models.Product

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especificar el nombre del producto "
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.ProdId = id
	err2 := bd.UpdateProduct(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar actualizar el registro de producto" + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, " Update OK"

}

func DeleteProduct(body string, User string, id int) (int, string) {
	fmt.Println("Inicializando funcion  router.DeleteProduct")

	if id == 0 {
		return 400, "Debe especificar ID del Producto a borrar  "
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err := bd.DeleteProduct(id)
	if err != nil {
		return 400, "Ocurrio un error al intentar borrar el producto" + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, " Delete OK"
}

func SelectProduct(body string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Inicializando funcion  router.SelectProduct")

	var t models.Product
	var page, pageSize int
	var orderType, orderField string

	param := request.QueryStringParameters

	page, _ = strconv.Atoi(param["page"])
	pageSize, _ = strconv.Atoi(param["pageSize"])
	orderType = param["orderType"]   //D= Desc, A o Nil= Asc
	orderField = param["orderField"] //I Id, T title, D Description, F Created At
	//P price, C CategoryId, S Stock

	if !strings.Contains("ITDFPCS", orderField) {
		orderField = ""
	}

	var choice string
	if len(param["prodId"]) > 0 {
		choice = "P"
		t.ProdId, _ = strconv.Atoi(param["prodId"])
	}

	if len(param["search"]) > 0 {
		choice = "S"
		t.ProdSearch, _ = param["search"]
	}

	if len(param["categId"]) > 0 {
		choice = "C"
		t.ProdCategId, _ = strconv.Atoi(param["categId"])
	}

	if len(param["slug"]) > 0 {
		choice = "U"
		t.ProdPath, _ = param["slug"]
	}

	if len(param["slugCateg"]) > 0 {
		choice = "K"
		t.ProdCategPath, _ = param["slugCateg"]
	}

	fmt.Println(param)

	listResult, err2 := bd.SelectProduct(t, choice, page, pageSize, orderType, orderField)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar capturar Resultados de tipo/s > " + choice + "' en productos > " + err2.Error()
	}

	Product, err3 := json.Marshal(listResult)
	if err3 != nil {
		return 400, "Ocurrio un error al intentar convertir en JSON  la busqueda de productos/s > " + err3.Error()
	}

	return 200, string(Product)

}

func UpdateStock(body string, User string, id int) (int, string) {
	fmt.Println("Inicializando funcion  router.UpdateStock")

	var t models.Product

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	if len(t.ProdTitle) == 0 {
		return 400, "Debe especificar el nombre del producto "
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.ProdId = id
	err2 := bd.UpdateStock(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar actualizar el Stock" + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, " Update Stock OK"

}
