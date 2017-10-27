package db

//Access to DB
import (
	"log"
	"reflect"

	mgo "gopkg.in/mgo.v2"
)

//GetDbSession retorna la session de la BD
func GetDbSession() *mgo.Session {
	//creacion conexion localmente con ReplicaSet, 1Primmary, 2 Slaves, con driver mgo
	session, err := mgo.Dial("localhost" /*:27017,localhost:27018,localhost:27019*/)
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

	var c *mgo.Collection
	valueSlice := reflect.New(
		reflect.MakeSlice(
			reflect.SliceOf(
				reflect.TypeOf(obj),
			), 0, 0,
		).Type(),
	)

	c = session.DB("lushflydb").C(entitieType)
	err := c.Find(nil).Sort("-start").All(valueSlice.Interface()) //es opcional el sort
	if err != nil {
		panic(err)
	}
	return valueSlice.Interface()
}

// //GetObjs Metodo que retorna objetos genericos
// func GetObjs(entitieType string, objs *reflect.Value) {
// 	session := GetDbSession()
// 	defer session.Close()

// 	var c *mgo.Collection

// 	c = session.DB("lushflydb").C(entitieType)
// 	err := c.Find(nil).Sort("-start").All(objs.Interface()) //es opcional el sort
// 	if err != nil {
// 		panic(err)
// 	}
// }

//InsertObj Metodo de insercion generico
func InsertObj(obj interface{}, entitieType string) {
	session := GetDbSession()
	defer session.Close()

	var c *mgo.Collection

	c = session.DB("lushflydb").C(entitieType)
	err := c.Insert(obj)
	if err != nil {
		log.Fatal(err)
	}
}
