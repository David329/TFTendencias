package db

//Access to DB
import (
	"log"
	"reflect"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	connections = "localhost" /*:27017,localhost:27018,localhost:27019*/
	dbname      = "lushflydb"
)

//GetDbSession retorna la session de la BD
func GetDbSession() *mgo.Session {
	//creacion conexion localmente con ReplicaSet, 1Primmary, 2 Slaves, con driver mgo
	session, err := mgo.Dial(connections)
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

//GetObjs Metodo que retorna objetos genericos;;//PADAVID: ->creo q en ves de pasar ese string podriamos gettypestring, lo malo q tiraria User y no Users; pa la proxxx
func GetObjs(entitieType string, obj interface{}) interface{} {
	session := GetDbSession()
	defer session.Close()

	valueSlice := reflect.New(reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(obj)), 0, 0).Type())

	err := session.DB(dbname).C(entitieType).Find(nil).Sort("-start").All(valueSlice.Interface()) //es opcional el sort
	if err != nil {
		panic(err)
	}
	return valueSlice.Interface()
}

//GetObjsByID ...
func GetObjsByID(entitieType string, id string, obj interface{}) {
	session := GetDbSession()
	defer session.Close()

	err := session.DB(dbname).C(entitieType).FindId(bson.ObjectIdHex(id)).One(obj)

	if err != nil {
		log.Fatal(err)
	}
}

//InsertObj Metodo de insercion generico
func InsertObj(entitieType string, obj interface{}) {
	session := GetDbSession()
	defer session.Close()

	err := session.DB(dbname).C(entitieType).Insert(obj)

	if err != nil {
		log.Fatal(err)
	}
}

//UpdateObjByID ...
func UpdateObjByID(entitieType string, id string, obj interface{}) {
	session := GetDbSession()
	defer session.Close()

	err := session.DB(dbname).C(entitieType).UpdateId(bson.ObjectIdHex(id), obj)

	if err != nil {
		log.Fatal(err)
	}
}

//DeleteObjByID ...
func DeleteObjByID(entitieType string, id string) {
	session := GetDbSession()
	defer session.Close()

	err := session.DB(dbname).C(entitieType).RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		log.Fatal(err)
	}
}
