package main

import (
	"log"

	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"//bson para formato como select where balblala, chekar la documentacion
)
////DOCUMENTATION
//https://godoc.org/gopkg.in/mgo.v2
//http://www.jancarloviray.com/blog/go-mongodb-simple-example/
type Payment struct{
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
//no intentar con la red de la upc, no abrexd
func main() {
	session, err := mgo.Dial("mongodb://userdb:passworddb@ds149324.mlab.com:49324/lushflydb")
	if err != nil {
			panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)


		// //Pa' insertar
// objAInsertar:= Users{
// 	FirstName: "ffffcsdsd",
// 	LastName: "xtuki",
// 	PassportType: "DNI",
// 	PassportNumber: "434",
// 	Email: "david@gmail.com",
// 	Password: "iksdi",
// 	PersonalCard: Payment{
// 		Card: "Visa",
// 		CardNumber: "2323",
// 		CSC: "2232323",
// 		Total: 2332,
// 		ExpirationDate: "23/21/1923",
// 	},
// }
	// c:= session.DB("lushflydb").C("Users")
	// // err = c.Insert(&Users{Name:"asdsd"})
	// err=c.Insert(objAInsertar)
	// if err != nil {
	//  log.Fatal(err)
	// }
	// log.Print(objAInsertar.FirstName);

	//Pa' Obtener
	var users []Users
	c:=	session.DB("lushflydb").C("Users")
	err = c.Find(nil).Sort("-start").All(&users)//es opcional el sort
	if err != nil {
		panic(err)
	}

	log.Print(users[0].LastName)//pera
}
