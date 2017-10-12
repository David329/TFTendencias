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
	response, err := http.Get(BackendURL + "/users")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &obj)

	testTemplate, _ := template.ParseFiles("./View/users.gohtml")

	wr.Header().Set("Content-Type", "text/html")

	err = testTemplate.Execute(wr, obj)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
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
