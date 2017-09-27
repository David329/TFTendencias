package main

import (
	"log"

	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"//bson para formato como select where balblala, chekar la documentacion
)
////DOCUMENTATION
//https://godoc.org/gopkg.in/mgo.v2
//http://www.jancarloviray.com/blog/go-mongodb-simple-example/
type Users struct {
	Name string
}

func main() {
	//Agregar las credenciales, con el usuario asignado a la BD lushflydb
	session, err := mgo.Dial("mongodb://usuario:contrasena@ds149324.mlab.com:49324/lushflydb")
	if err != nil {
			panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// //Pa' insertar
	// c:= session.DB("lushflydb").C("Users")
	// err = c.Insert(&Users{Name:"asdsd"})
	// if err != nil {
	//  log.Fatal(err)
	// }
	
	//Pa' Obtener
	var users []Users
	c:=	session.DB("lushflydb").C("Users")
	err = c.Find(nil).Sort("-start").All(&users)//es opcional el sort
	if err != nil {
		panic(err)
	}
	log.Print(users[1].Name)//pera
}
