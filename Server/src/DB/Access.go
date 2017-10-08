package db

//Access to DB
import (
	"encoding/json"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

//GetDbSession retorna la session de la BD
func GetDbSession() *mgo.Session {
	//creacion conexion con mlab, con driver mgo
	session, err := mgo.Dial("mongodb://userdb:passworddb@ds149324.mlab.com:49324/lushflydb")
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

//SendResCloseSession Envia Respuesta y Cierra la conexion
func SendResCloseSession(message string, session *mgo.Session, wr http.ResponseWriter) {
	//formato de envio
	wr.Header().Set("Content-Type", "application/json")

	//estado web
	wr.WriteHeader(http.StatusOK)

	//envio del json a la ruta
	json.NewEncoder(wr).Encode(message)
	session.Close()
}
