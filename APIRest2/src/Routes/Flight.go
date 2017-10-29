//Package routes allow methods for Model Flight
package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	DB "../DB"
	Entities "../Entities"

	"github.com/julienschmidt/httprouter"
)

//GetAllFlight Return All Objects
func GetAllFlight(wr http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	//Return through pointer and save in obj
	var obj []interface{}
	DB.GetObjs("Flights", &obj)

	//Response ok or error
	response(&wr, &obj[0])
}

//GetFlightByID Return object ByID
func GetFlightByID(wr http.ResponseWriter, _ *http.Request, ps httprouter.Params) {

	//Return through pointer and save in obj
	var obj interface{}
	DB.GetObjsByID("Flights", ps.ByName("id"), &obj)

	//Response ok or error
	response(&wr, &obj)
}

//PostFlight Insert a new Object
func PostFlight(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//New Obj like Entities.Flight
	var obj interface{} = new(Entities.Flight)

	//Read Body of Form, then convert json binary to Struct previously defined
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &obj)

	//Insert obj, then return through pointer
	DB.InsertObj("Flights", &obj)

	//Response ok or error
	response(&wr, &obj)
}

//PutFlightByID Update a Object
func PutFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	//New Obj like Entities.Flight
	var obj interface{} = new(Entities.Flight)

	//Read Body of Form, then convert json binary to Struct previously defined
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &obj)

	//Update obj, then return through pointer
	DB.UpdateObjByID("Flights", ps.ByName("id"), &obj)

	//Response ok or error
	response(&wr, &obj)
}

//DeleteFlightByID Delete object by ID
func DeleteFlightByID(wr http.ResponseWriter, _ *http.Request, ps httprouter.Params) {

	//Delete object depending Model's ID of url, then return through pointer
	var obj interface{} = ps.ByName("id")
	DB.DeleteObjByID("Flights", &obj)

	//Response ok or error
	response(&wr, &obj)
}
