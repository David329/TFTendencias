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
    Entities "./Entities"

    "github.com/gorilla/pat"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson" //bson para formato como select where balblala, chekar la documentacion
)

func getDbSession()(*mgo.Session){
    //creacion conexion con mlab, con driver mgo
    session, err:= mgo.Dial("mongodb://userdb:passdb@ds149324.mlab.com:49324/lushflydb")
    if err != nil {
        log.Print(err)
    }
    session.SetMode(mgo.Monotonic, true)
    return session
}
func sendResCloseSession(message string, session mgo.Session, wr http.ResponseWriter)(){
    //formato de envio
    wr.Header().Set("Content-Type", "application/json")

    //estado web
    wr.WriteHeader(http.StatusOK)

    //envio del json a la ruta
    json.NewEncoder(wr).Encode(message)
    session.Close()
}

//FLIGHTS
func postFlights(wr http.ResponseWriter, req *http.Request) {

    session:= getDbSession()

    //obtener el json y lo guardo en body
    var obj Entities.Flights
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

    sendResCloseSession("Objeto Insertado",*session,wr)
}
//IMPORTANTE ESTE UPDATE NOS PODRIA SERVIR PARA ATRIBUTOS DEL VUELO, NO PARA ACTUALIZAR LOS ASIENTOS OK...
//PARA ACTUALIZAR LOS ASIENTOS DEBEMOS OBTENER EL USUARIO Y OBTENER EL OBJETO DEL VUELO Y LUEGO AGREGAR Y CONECTAR EN ASIENTOS DE FLIGHT
func putFlightsById(wr http.ResponseWriter, req *http.Request) {

    session:= getDbSession()

    //obtener el json y lo guardo en body
    var obj Entities.Flights
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

    sendResCloseSession("Actualizacion Exitosa",*session,wr)
}

//USERS
func getAllUsers(wr http.ResponseWriter, req * http.Request) {

    session:= getDbSession()

    //Pa' Obtener
    var users[] Entities.Users
    c:= session.DB("lushflydb").C("Users")
    err:= c.Find(nil).Sort("-start").All(&users) //es opcional el sort
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

    sendResCloseSession(string(data),*session,wr)
}
func deleteUsersById(wr http.ResponseWriter, req * http.Request) {

    session:= getDbSession()

    //obtener el id desde la url
	reqId:= req.URL.Query().Get(":id")

    //obtener solo los q tienen ese id
	c:= session.DB("lushflydb").C("Users")
    err:= c.RemoveId(bson.ObjectIdHex(reqId))
    
    if err != nil {
        panic(err)
    }

    sendResCloseSession("Eliminacion Exitosa",*session,wr)
}

func main() {
    //creacion de enrutador pat
    router:= pat.New()

    //llamada de metodos a enrutar
    router.Get("/users", getAllUsers)
    router.Delete("/users/{id}", deleteUsersById)
    router.Post("/flights", postFlights)
    router.Put("/flights/{id}", putFlightsById)

    //activar entutador, probar a;adirlo en otro archivo...
    http.Handle("/", router)

    //escuchando el servidor
    log.Print("Escuchando en 127.0.0.1:8000...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}