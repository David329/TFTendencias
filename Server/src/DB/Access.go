package db

//Access to DB
import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

//GetDbSession retorna la session de la BD
func GetDbSession() *mgo.Session {
	//creacion conexion con mlab, con driver mgo
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
