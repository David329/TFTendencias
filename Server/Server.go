////DOCUMENTATION
//https://godoc.org/gopkg.in/mgo.v2
//http://www.jancarloviray.com/blog/go-mongodb-simple-example/
//PREGUNTAR LAS CREDENCIALES DE USUARIO DE BASE DE DATOS
package main

import (
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"

    "github.com/gorilla/pat"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson" //bson para formato como select where balblala, chekar la documentacion
)
type Seat struct {
    UserID string
    Number string
    Type string
}
type Departure struct {
    Country string
    City string
    TD string //departure, tndria q ser un dateTime
    TA string //arrival, tndria q ser un dateTime
}
type Destination struct {
    Country string
    City string
    TD string //departure, tndria q ser un dateTime
    TA string //arrival, tndria q ser un dateTime	
}
type Flights struct {
    AirplaneModel string
    AirplaneNumber string
    Price float32
    Depart Departure
    Destin Destination
    Seats[] Seat
}
//asociar payment, usuario y reserva
type Payment struct {
    Card string
    CardNumber string
    CSC string
    Total float32
    ExpirationDate string //format=(dd/mm/yyyy)
}
type Users struct {
    FirstName string
    LastName string
    PassportType string
    PassportNumber string
    Email string
    Password string
    PersonalCard Payment
}

//FLIGHTS
func postFlights(wr http.ResponseWriter, req *http.Request) {

    //creacion conexion con mlab, con driver mgo
    session, err:= mgo.Dial("mongodb://userdb:passworddb@ds149324.mlab.com:49324/lushflydb")
    session.SetMode(mgo.Monotonic, true)

    //obtener el json y lo guardo en body
    var obj Flights
    body, err:= ioutil.ReadAll(req.Body)
    if err != nil {
        log.Print(err)
    }

    //parseo de json a flight, nose si parsea mas de 1 objeto..., seguro con un for o algo
    json.Unmarshal(body, &obj)

    //inserto en la bd
    c:= session.DB("lushflydb").C("Flights")
    err = c.Insert(obj)
    if err != nil {
        log.Fatal(err)
    }

    //respuesta
    wr.Header().Set("Content-Type", "application/json")
    wr.WriteHeader(http.StatusOK)
    json.NewEncoder(wr).Encode("Obj inserted...")
}
//IMPORTANTE ESTE UPDATE NOS PODRIA SERVIR PARA ATRIBUTOS DEL VUELO, NO PARA ACTUALIZAR LOS ASIENTOS OK...
//PARA ACTUALIZAR LOS ASIENTOS DEBEMOS OBTENER EL USUARIO Y OBTENER EL OBJETO DEL VUELO Y LUEGO AGREGAR Y CONECTAR EN ASIENTOS DE FLIGHT
func putFlights(wr http.ResponseWriter, req *http.Request) {

    //creacion conexion con mlab, con driver mgo
    session, err:= mgo.Dial("mongodb://userdb:passworddb@ds149324.mlab.com:49324/lushflydb")
    session.SetMode(mgo.Monotonic, true)

    //obtener el json y lo guardo en body
    var obj Flights
    body, err:= ioutil.ReadAll(req.Body)
    if err != nil {
        log.Print(err)
    }

    //parseo de json a flight, lo hacemos para actualizar los asientos
	json.Unmarshal(body, &obj)
	
    //obtener el id desde la url
	reqId:= req.URL.Query().Get(":id")

    //obtener solo los q tienen ese id
	c:= session.DB("lushflydb").C("Flights")
	err = c.UpdateId(bson.ObjectIdHex(reqId),obj)

	////tmb se puede modificar asi, solo atributos
    // err = c.UpdateId(bson.ObjectIdHex(reqId), bson.M {
    //         "$set": bson.M {
    //             "airplanemodel": "uptddsds"
    //         }
	// 	})

	//err= c.FindId(bson.ObjectIdHex(reqId)).Sort("-start").One(&obj)//asi obtenemos un objeto
    if err != nil {
        panic(err)
    }

    //respuesta
    wr.Header().Set("Content-Type", "application/json")
    wr.WriteHeader(http.StatusOK)
    json.NewEncoder(wr).Encode("Actualizacion Exitosa")
}

//USERS
func getAllUsers(wr http.ResponseWriter, req * http.Request) {

    //creacion conexion con mlab, con driver mgo
    session, err:= mgo.Dial("mongodb://userdb:passworddb@ds149324.mlab.com:49324/lushflydb")
    session.SetMode(mgo.Monotonic, true)

    //Pa' Obtener
    var users[] Users
    c:= session.DB("lushflydb").C("Users")
    err = c.Find(nil).Sort("-start").All(&users) //es opcional el sort
    if err != nil {
        panic(err)
    }

    //cerrramos sesion
    session.Close()

    //parseamos a json
    data, err:= json.Marshal(users)
    if err != nil {
        log.Print(err)
        return
    }

    //Formato Json a mostrar
    wr.Header().Set("Content-Type", "application/json")

    //estado web
    wr.WriteHeader(http.StatusOK)

    //envio del json a la ruta
    json.NewEncoder(wr).Encode(string(data))
}

func main() {
    //creacion de enrutador pat
    router:= pat.New()

    //llamada de metodos a enrutar
    router.Get("/users", getAllUsers)
    router.Post("/flights", postFlights)
    router.Put("/flights/{id}", putFlights)

    //activar entutador, probar a;adirlo en otro archivo...
    http.Handle("/", router)

    //escuchando el servidor
    log.Print("Escuchando en 127.0.0.1:8000...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}