package routes

//Restful - Flight
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	DB "../DB"
	Entities "../Entities"

	"gopkg.in/mgo.v2/bson"
)

//PostFlight Inserta un nuevo vuelo
func PostFlight(wr http.ResponseWriter, req *http.Request) {

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

	DB.SendResCloseSession("Objeto Insertado", session, wr)
}

//PutFlightByID Actualiza un Documento Flight
func PutFlightByID(wr http.ResponseWriter, req *http.Request) {

	session := DB.GetDbSession()

	//obtener el json y lo guardo en body
	var obj Entities.Flight
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a flight, lo hacemos para actualizar los asientos
	json.Unmarshal(body, &obj)

	//obtener el id desde la url
	reqID := req.URL.Query().Get(":id")

	//obtener solo los q tienen ese id
	c := session.DB("lushflydb").C("Flights")
	err = c.UpdateId(bson.ObjectIdHex(reqID), obj)

	////tmb se puede modificar asi, solo atributos
	// err = c.UpdateId(bson.ObjectIdHex(reqID), bson.M {
	//         "$set": bson.M {
	//             "airplanemodel": "uptddsds"
	//         }
	// 	})

	//err= c.FindId(bson.ObjectIdHex(reqID)).Sort("-start").One(&obj)//asi obtenemos un objeto
	if err != nil {
		panic(err)
	}

	DB.SendResCloseSession("Actualizacion Exitosa", session, wr)
}

//IMPORTANTE ESTE UPDATE NOS PODRIA SERVIR PARA ATRIBUTOS DEL VUELO, NO PARA ACTUALIZAR LOS ASIENTOS OK...
//PARA ACTUALIZAR LOS ASIENTOS DEBEMOS OBTENER EL USUARIO Y OBTENER EL OBJETO DEL VUELO Y LUEGO AGREGAR Y CONECTAR EN ASIENTOS DE FLIGHT
