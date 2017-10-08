package main

import (
	"fmt"
	"net/http"

	Routes "./Routes"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/users", Routes.Lushoreq)

	fmt.Println("Servidor Corriendo en puerto 8080")

	http.ListenAndServe(":8080", router)

}
