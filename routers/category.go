package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/livingstone23/gambit/bd"
	"github.com/livingstone23/gambit/models"
)

func InsertCategory(body string, User string) (int, string) {
	fmt.Println("Inicializando funcion  router.InsertCategory")

	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "Debe especificar el nombre (title) de la categoria"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el registro de categoria" + t.CategName + " > " + err2.Error()
	}

	return 200, "{CategID: " + strconv.Itoa(int(result)) + "}"
}
