package routes

//Restful - User
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	DB "../DB"
	Entities "../Entities"
	"github.com/julienschmidt/httprouter"

	"gopkg.in/mgo.v2/bson"
)

//Get-Post-Put-Delete

//GetAllUser Envia todos los usuarios, formato->JSON
func GetAllUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

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

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(users)

}

//PostUser Inserta un nuevo vuelo
func PostUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	session := DB.GetDbSession()

	//obtener el json y lo guardo en body
	var obj Entities.User
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a User, nose si parsea mas de 1 objeto..., seguro con un for o algo
	json.Unmarshal(body, &obj)

	//inserto en la bd
	c := session.DB("lushflydb").C("Users")

	err = c.Insert(obj)
	if err != nil {
		log.Fatal(err)
	}

	//cerrramos sesion
	session.Close()

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Insertado")
}

//PutUserByID Actualiza un Documento User
func PutUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	session := DB.GetDbSession()

	//obtener el json y lo guardo en body
	var obj Entities.User
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a user
	json.Unmarshal(body, &obj)

	//obtener el id desde la url
	reqID := ps.ByName("id")

	//obtener solo los q tienen ese id
	c := session.DB("lushflydb").C("Users")

	err = c.UpdateId(bson.ObjectIdHex(reqID), obj)
	if err != nil {
		panic(err)
	}

	//cerrramos sesion
	session.Close()

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Actualizado")
}

//DeleteUserByID Elimina un usuario por ID, formato->JSON
func DeleteUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	session := DB.GetDbSession()

	//obtener el id desde la url
	reqID := ps.ByName("id")

	//obtener solo los q tienen ese id
	c := session.DB("lushflydb").C("Users")

	err := c.RemoveId(bson.ObjectIdHex(reqID))

	if err != nil {
		panic(err)
	}

	//cerrramos sesion
	session.Close()

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Eliminado")
}
