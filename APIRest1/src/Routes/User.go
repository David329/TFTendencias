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

//GetAllUserByLastName Metodo creado para probar el metodo getByQuery, la consulta debe estar en minuscula
func GetAllUserByLastName(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//Obtener Todos los usuarios por el metodo Generico
	var users []Entities.User
	users = *DB.GetObjsByQuery("Users", Entities.User{}, bson.M{"lastname": "Silvaxxxx"}).(*[]Entities.User)

	// //Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(users)

}

//GetAllUser Envia todos los usuarios, formato->JSON
func GetAllUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//Obtener Todos los usuarios por el metodo Generico
	var users []Entities.User
	users = *DB.GetObjs("Users", Entities.User{}).(*[]Entities.User)

	// //Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(users)

}

//GetUserByID Envia El usuario por ID, formato->JSON
func GetUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	user := &Entities.User{}

	DB.GetObjsByID("Users", ps.ByName("id"), &user)

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(user)
}

//PostUser Inserta un nuevo vuelo
func PostUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//obtener el json y lo guardo en body
	var obj Entities.User
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a User, nose si parsea mas de 1 objeto..., seguro con un for o algo
	json.Unmarshal(body, &obj)

	//Insercion en metodo generico
	DB.InsertObj("Users", obj)

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Insertado")
}

//PutUserByID Actualiza un Documento User
func PutUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	//obtener el json y lo guardo en body
	var obj Entities.User
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a user
	json.Unmarshal(body, &obj)
	DB.UpdateObjByID("Users", ps.ByName("id"), obj)

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Actualizado")
}

//DeleteUserByID Elimina un usuario por ID, formato->JSON
func DeleteUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	DB.DeleteObjByID("Users", ps.ByName("id"))

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Eliminado")
}
