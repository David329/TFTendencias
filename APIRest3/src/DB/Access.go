//Package db allow connection with DB
package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	connections = "localhost" /*:27017,localhost:27018,localhost:27019*/
	dbName      = "lushflydb"
	errorStatus = "error"
	okStatus    = "ok"
)

//getDbSession return session of DB
func getDbSession() *mgo.Session {

	//Create connection local with ReplicaSet, 1Primmary, 2 Slaves with driver mgo
	session, err := mgo.Dial(connections)
	if err != nil {
		log.Println(err)
	}

	//SetMode of behaviour of DB
	session.SetMode(mgo.Monotonic, true)
	return session
}

//getInOneObject transform array to one object(interface)
func getInOneObject(obj *[]interface{}) {

	//Aux for storage a array of interfaces
	var aux interface{}
	aux = *obj

	//Reset the length of bytes, then append the array in one interface
	(*obj) = *new([]interface{})
	(*obj) = append(*obj, aux)
}

//GetObjsByQuery return object/s given a query bson
func GetObjsByQuery(entitieType string, obj *[]interface{}, query interface{}) {

	session := getDbSession()
	defer session.Close()

	//Find all object/s where condition is correct, then save in array of interfaces
	err := session.DB(dbName).C(entitieType).Find(query).Sort("-start").All(obj)

	//Handle Error, reuse interface for response
	if err != nil {
		log.Println(err)
		(*obj) = *new([]interface{})
		(*obj) = append(*obj, errorStatus)
		return
	}

	getInOneObject(obj)
}

//GetObjs return objects depending on type of model;									//PADAVID: ->creo q en ves de pasar ese string podriamos gettypestring, lo malo q tiraria User y no Users; pa la proxxx
func GetObjs(entitieType string, obj *[]interface{}) {

	session := getDbSession()
	defer session.Close()

	//Find all objects, then save in array of interfaces
	err := session.DB(dbName).C(entitieType).Find(nil).Sort("-start").All(obj)

	//Handle Error, reuse interface for response
	if err != nil {
		log.Println(err)
		(*obj) = *new([]interface{})
		(*obj) = append(*obj, errorStatus)
		return
	}

	getInOneObject(obj)
}

//GetObjsByID return object depending on ID and type of model
func GetObjsByID(entitieType string, id string, obj *interface{}) {

	session := getDbSession()
	defer session.Close()

	//Find object, then save in obj pointer
	err := session.DB(dbName).C(entitieType).FindId(bson.ObjectIdHex(id)).One(obj)

	//Handle Error, reuse interface for response
	if err != nil {
		log.Println(err)
		(*obj) = errorStatus
		return
	}
}

//InsertObj save object in specific model
func InsertObj(entitieType string, obj *interface{}) {

	session := getDbSession()
	defer session.Close()

	//Insert object
	err := session.DB(dbName).C(entitieType).Insert(obj)

	//Handle Error, reuse interface for response
	if err != nil {
		log.Println(err)
		(*obj) = errorStatus
		return
	}
	(*obj) = okStatus
}

//UpdateObjByID update object in specific model given ID
func UpdateObjByID(entitieType string, id string, obj *interface{}) {

	session := getDbSession()
	defer session.Close()

	//Update object given ID
	err := session.DB(dbName).C(entitieType).UpdateId(bson.ObjectIdHex(id), obj)

	//Handle Error, reuse interface for response
	if err != nil {
		log.Println(err)
		(*obj) = errorStatus
		return
	}
	(*obj) = okStatus
}

//DeleteObjByID delete object in specific model given ID
func DeleteObjByID(entitieType string, id *interface{}) {

	session := getDbSession()
	defer session.Close()

	//Delete object given ID
	err := session.DB(dbName).C(entitieType).RemoveId(bson.ObjectIdHex((*id).(string)))

	//Handle Error, reuse interface for response
	if err != nil {
		log.Println(err)
		(*id) = errorStatus
		return
	}
	(*id) = okStatus
}
