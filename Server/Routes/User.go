package routes

//Restful - User
import (
	"encoding/json"
	"log"
	"net/http"

	DB "../DB"
	Entities "../Entities"

	"gopkg.in/mgo.v2/bson"
)

//GetAllUser Envia todos los usuarios, formato->JSON
func GetAllUser(wr http.ResponseWriter, req *http.Request) {

	session := DB.GetDbSession() //en mayusculas pa q sea publico

	//Pa' Obtener
	var users []Entities.User
	c := session.DB("lushflydb").C("Users")
	err := c.Find(nil).Sort("-start").All(&users) //es opcional el sort
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

	DB.SendResCloseSession(string(data), session, wr)
}

//DeleteUserByID Elimina un usuario por ID, formato->JSON
func DeleteUserByID(wr http.ResponseWriter, req *http.Request) {

	session := DB.GetDbSession()

	//obtener el id desde la url
	reqID := req.URL.Query().Get(":id")

	//obtener solo los q tienen ese id
	c := session.DB("lushflydb").C("Users")
	err := c.RemoveId(bson.ObjectIdHex(reqID))

	if err != nil {
		panic(err)
	}

	DB.SendResCloseSession("Eliminacion Exitosa", session, wr)
}
