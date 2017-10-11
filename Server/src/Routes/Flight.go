package routes

//Restful - Flight
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

//GetAllFlight Envia todos los vuelos, formato->JSON
func GetAllFlight(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	session := DB.GetDbSession() //en mayusculas pa q sea publico

	//Pa' Obtener
	var flights []Entities.Flight
	c := session.DB("lushflydb").C("Flights")

	err := c.Find(nil).Sort("-start").All(&flights) //es opcional el sort
	if err != nil {
		panic(err)
	}

	//cerrramos sesion
	session.Close()

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(flights)
}

//PostFlight Inserta un nuevo vuelo
func PostFlight(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	session := DB.GetDbSession()

	//obtener el json y lo guardo en body
	var obj Entities.Flight
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a flight, nose si parsea mas de 1 objeto..., seguro con un for o algo
	json.Unmarshal(body, &obj)

	//inserto en la bd
	c := session.DB("lushflydb").C("Flights")

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

//PutFlightByID Actualiza un Documento Flight
func PutFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	session := DB.GetDbSession()

	//obtener el json y lo guardo en body
	var obj Entities.Flight
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a flight
	json.Unmarshal(body, &obj)

	//obtener el id desde la url
	reqID := ps.ByName("id")

	//obtener solo los q tienen ese id
	c := session.DB("lushflydb").C("Flights")

	err = c.UpdateId(bson.ObjectIdHex(reqID), obj)
	if err != nil {
		panic(err)
	}
	////tmb se puede modificar asi, solo atributos
	// err = c.UpdateId(bson.ObjectIdHex(reqID), bson.M {
	//         "$set": bson.M {
	//             "airplanemodel": "uptddsds"
	//         }
	// 	})

	//err= c.FindId(bson.ObjectIdHex(reqID)).Sort("-start").One(&obj)//asi obtenemos un objeto

	//cerrramos sesion
	session.Close()

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Actualizado")
}

//DeleteFlightByID Elimina un usuario por ID, formato->JSON
func DeleteFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	session := DB.GetDbSession()

	//obtener el id desde la url
	reqID := ps.ByName("id")

	//obtener solo los q tienen ese id
	c := session.DB("lushflydb").C("Flights")

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
