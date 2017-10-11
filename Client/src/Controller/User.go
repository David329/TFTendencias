package controller

//Restful - User
import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	Models "../Model"

	"github.com/julienschmidt/httprouter"
)

//Get-Post-Put-Delete

const (
	//BackendURL asd
	BackendURL = "http://localhost:8000/"
)

//GetAllUser Envia todos los usuarios, formato->JSON
func GetAllUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var obj []Models.User
	response, err := http.Get("http://localhost:8000/users")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &obj)

	view := "users.html"

	log.Print(obj)

	t, _ := template.ParseFiles(view /* + ".html"*/)

	log.Print(t)

	t.Execute(wr, obj[0])

}

// //PostUser Inserta un nuevo vuelo
// func PostUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

// }

// //PutUserByID Actualiza un Documento User
// func PutUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

// }

// //DeleteUserByID Elimina un usuario por ID, formato->JSON
// func DeleteUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

// }
