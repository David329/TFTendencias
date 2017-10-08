package routes

import (
	"encoding/json"
	"log"
	"net/http"

	Entities "../Entities"
	mgo "gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
)

//Lushoreq prueba
func Lushoreq(wr http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session, err := mgo.Dial("mongodb://guest:guest@ds149324.mlab.com:49324/lushflydb")
	if err != nil {
		log.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)

	//Pa' Obtener
	var users []Entities.User
	c := session.DB("lushflydb").C("Users")
	err = c.Find(nil).Sort("-start").All(&users) //es opcional el sort
	if err != nil {
		panic(err)
	}

	//cerrramos sesion
	session.Close()

	//parseamos a json
	data, err := json.Marshal(users)
	if err != nil {
		log.Print(err)
		return
	}

	wr.Header().Set("Content-Type", "application/json")

	//estado web
	wr.WriteHeader(http.StatusOK)

	//envio del json a la ruta
	json.NewEncoder(wr).Encode(string(data))
	session.Close()
}
