//Package routes allows methods for Model User
package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	DB "../DB"
	Entities "../Entities"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

//response
func response(wr *http.ResponseWriter, obj *interface{}) {

	//response in json format
	(*wr).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*wr).Encode(*obj)
}

//GetAllUserByLastName allow test GenericMethod -> GetObjsByQuery
func GetAllUserByLastName(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//Get All Users where lastname is X by GenericMethod
	// var users []Entities.User
	var users interface{}
	users = *DB.GetObjsByQuery("Users", Entities.User{}, bson.M{"lastname": "Silvaxxxx"}).(*[]Entities.User)

	//Response
	//	wr.Header().Set("Content-Type", "application/json")
	//	json.NewEncoder(wr).Encode(users)
	response(&wr, &users)
}

//GetAllUser Return All Users -> JSON
func GetAllUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//Obtener Todos los usuarios por el metodo Generico
	// var users []Entities.User
	var users interface{}
	users = *DB.GetObjs("Users", Entities.User{}).(*[]Entities.User)

	//Response
	//	wr.Header().Set("Content-Type", "application/json")
	//	json.NewEncoder(wr).Encode(users)
	response(&wr, &users)

}

//GetUserByID Envia El usuario por ID, formato->JSON
func GetUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	user := &Entities.User{}

	DB.GetObjsByID("Users", ps.ByName("id"), &user)

	//Response
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(user)
}

//PostUser Inserta un nuevo vuelo
func PostUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//obtener el json y lo guardo en body
	var obj Entities.User
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	//parseo de json a User, nose si parsea mas de 1 objeto..., seguro con un for o algo
	json.Unmarshal(body, &obj)

	//Insercion en metodo generico
	DB.InsertObj("Users", obj)

	//Response
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Insertado")
}

//PutUserByID Actualiza un Documento User
func PutUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	//obtener el json y lo guardo en body
	var obj Entities.User
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	//parseo de json a user
	json.Unmarshal(body, &obj)
	DB.UpdateObjByID("Users", ps.ByName("id"), obj)

	//Response
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Actualizado")
}

//DeleteUserByID Elimina un usuario por ID, formato->JSON
func DeleteUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	DB.DeleteObjByID("Users", ps.ByName("id"))

	//Response
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Eliminado")
}
