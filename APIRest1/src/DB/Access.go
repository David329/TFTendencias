package db

//Access to DB
import (
	"log"

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
