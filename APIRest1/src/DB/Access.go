//Package db allows connection with DB
package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	connections = "localhost" /*:27017,localhost:27018,localhost:27019*/
	dbname      = "lushflydb"
)

//getDbSession return session of DB
func getDbSession() *mgo.Session {

	//Create connection local with ReplicaSet, 1Primmary, 2 Slaves with driver mgo
	session, err := mgo.Dial(connections)
	if err != nil {
		log.Print(err)
	}

	//SetMode of behaviour of DB
	session.SetMode(mgo.Monotonic, true)
	return session
}

//getInOneObject ...
func getInOneObject(obj *[]interface{}) {
	var aux interface{}
	aux = *obj
	(*obj) = *new([]interface{})
	(*obj) = append(*obj, aux)
}

//GetObjsByQuery return object/s given a query bson;;//PADAVID: ->creo q en ves de pasar ese string podriamos gettypestring, lo malo q tiraria User y no Users; pa la proxxx
func GetObjsByQuery(entitieType string, obj *[]interface{}, query interface{}) {

	//GetSession
	session := getDbSession()
	defer session.Close()

	err := session.DB(dbname).C(entitieType).Find(query).Sort("-start").All(obj)
	if err != nil {
		panic(err)
	}
	getInOneObject(obj)
}

//GetObjs Metodo que retorna objetos genericos;;//PADAVID: ->creo q en ves de pasar ese string podriamos gettypestring, lo malo q tiraria User y no Users; pa la proxxx
func GetObjs(entitieType string, obj *[]interface{}) {
	session := getDbSession()
	defer session.Close()

	err := session.DB(dbname).C(entitieType).Find(nil).Sort("-start").All(obj)
	if err != nil {
		panic(err)
	}

	getInOneObject(obj)
}

//GetObjsByID ...
func GetObjsByID(entitieType string, id string, obj *interface{}) {
	session := getDbSession()
	defer session.Close()

	err := session.DB(dbname).C(entitieType).FindId(bson.ObjectIdHex(id)).One(obj)

	if err != nil {
		log.Fatal(err)
	}
}

//InsertObj Metodo de insercion generico
func InsertObj(entitieType string, obj *interface{}) {
	session := getDbSession()
	defer session.Close()

	err := session.DB(dbname).C(entitieType).Insert(obj)

	if err != nil {
		log.Fatal(err)
	}
}

//UpdateObjByID ...
func UpdateObjByID(entitieType string, id string, obj *interface{}) {
	session := getDbSession()
	defer session.Close()

	err := session.DB(dbname).C(entitieType).UpdateId(bson.ObjectIdHex(id), obj)

	if err != nil {
		log.Fatal(err)
	}
}

//DeleteObjByID ...
func DeleteObjByID(entitieType string, id string) {
	session := getDbSession()
	defer session.Close()

	err := session.DB(dbname).C(entitieType).RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		log.Fatal(err)
	}
}
