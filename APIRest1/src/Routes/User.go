package routes

//Restful - User
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	DB "../DB"
	Entities "../Entities"
	"github.com/julienschmidt/httprouter"

	"gopkg.in/mgo.v2/bson"
)

//Get-Post-Put-Delete
func ToIntf(s interface{}) []interface{} {
	v := reflect.ValueOf(s)
	// There is no need to check, we want to panic if it's not slice or array
	intf := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		intf[i] = v.Index(i).Interface()
	}
	return intf
}

//GetAllUser Envia todos los usuarios, formato->JSON
func GetAllUser(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//Obtener Todos los usuarios por el metodo Generico
	var users []Entities.User
	users = *DB.GetObjs("Users", Entities.User{}).(*[]Entities.User)

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(users)

}

//GetUserByID Envia El usuario por ID, formato->JSON
func GetUserByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	session := DB.GetDbSession()
	var user Entities.User

	c := session.DB("lushflydb").C("Users")

	c.FindId(bson.ObjectIdHex(ps.ByName("id"))).One(&user)

	//cerrramos sesion
	session.Close()

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
	DB.InsertObj(obj, "Users")

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
