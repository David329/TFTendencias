//Package routes allows methods for Model User
package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	DB "../DB"
	Entities "../Entities"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

//response through pointers, this method works for all Routes's Models
func response(wr *http.ResponseWriter, obj *interface{}) {

	//response in json format
	(*wr).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*wr).Encode(*obj)
}

//GetAllUserByLastName allow test GenericMethod -> GetObjsByQuery
func GetAllUserByLastName(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var obj []interface{}
	DB.GetObjsByQuery("Users", &obj, bson.M{"lastname": "Silvaxxxx"})

	response(&wr, &obj[0])
}

//GetAllUser Return All Users
func GetAllUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var obj []interface{}
	DB.GetObjs("Users", &obj)

	response(&wr, &obj[0])
}

//GetUserByID Return User ByID
func GetUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var obj interface{}
	DB.GetObjsByID("Users", ps.ByName("id"), &obj)

	response(&wr, &obj)
}

//PostUser Inserta un nuevo vuelo
func PostUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var obj interface{} = new(Entities.User)

	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &obj)

	DB.InsertObj("Users", &obj)

	obj = "ok"
	response(&wr, &obj)
}

//PutUserByID Actualiza un Documento User
func PutUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var obj interface{} = new(Entities.User)
	body, _ := ioutil.ReadAll(req.Body)

	json.Unmarshal(body, &obj)
	DB.UpdateObjByID("Users", ps.ByName("id"), &obj)

	obj = "ok"
	response(&wr, &obj)
}

//DeleteUserByID Elimina un usuario por ID, formato->JSON
func DeleteUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var obj interface{} = "ok"
	DB.DeleteObjByID("Users", ps.ByName("id"))

	response(&wr, &obj)
}
