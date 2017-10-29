//Package routes allow methods for Model User
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
func GetAllUserByLastName(wr http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	//Return through pointer and save in obj
	var obj []interface{}
	DB.GetObjsByQuery("Users", &obj, bson.M{"lastname": "Silvaxxxx"})

	//Response ok or error
	response(&wr, &obj[0])
}

//GetAllUser Return All Objects
func GetAllUser(wr http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	//Return through pointer and save in obj
	var obj []interface{}
	DB.GetObjs("Users", &obj)

	//Response ok or error
	response(&wr, &obj[0])
}

//GetUserByID Return object ByID
func GetUserByID(wr http.ResponseWriter, _ *http.Request, ps httprouter.Params) {

	//Return through pointer and save in obj
	var obj interface{}
	DB.GetObjsByID("Users", ps.ByName("id"), &obj)

	//Response ok or error
	response(&wr, &obj)
}

//PostUser Insert a new Object
func PostUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//New Obj like Entities.User
	var obj interface{} = new(Entities.User)

	//Read Body of Form, then convert json binary to Struct previously defined
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &obj)

	//Insert obj, then return through pointer
	DB.InsertObj("Users", &obj)

	//Response ok or error
	response(&wr, &obj)
}

//PutUserByID Update a Object
func PutUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	//New Obj like Entities.User
	var obj interface{} = new(Entities.User)

	//Read Body of Form, then convert json binary to Struct previously defined
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &obj)

	//Update obj, then return through pointer
	DB.UpdateObjByID("Users", ps.ByName("id"), &obj)

	//Response ok or error
	response(&wr, &obj)
}

//DeleteUserByID Delete object by ID
func DeleteUserByID(wr http.ResponseWriter, _ *http.Request, ps httprouter.Params) {

	//Delete object depending Model's ID of url, then return through pointer
	var obj interface{} = ps.ByName("id")
	DB.DeleteObjByID("Users", &obj)

	//Response ok or error
	response(&wr, &obj)
}
