////DOCUMENTATION
//https://godoc.org/gopkg.in/mgo.v2
//http://www.jancarloviray.com/blog/go-mongodb-simple-example/
//PREGUNTAR LAS CREDENCIALES DE USUARIO DE BASE DE DATOS -> ./DB/Access.go
package main

//MEJORAR:
//  1. Pasar los parametros como direccion de memoria, para eliminar la sesion y variables
//  2. Por ahora estamos enviando desde ./Test/SendDotNet/Send#, esto es solo de prueba, xq debe ser enviado desde el ClienteWeb.
import (
	"log"
	"net/http"
	"sync"

	Controller "./Controller"

	"github.com/julienschmidt/httprouter"
	//	"github.com/streadway/amqp"
)

//Parallelize funcion personalizada para que pueda correr mas de 1 gorutina
func Parallelize(functions ...func()) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))

	defer waitGroup.Wait()

	for _, function := range functions {
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(function)
	}
}

//**-MAIN-**
func main() {

	listenFirstChannel := func() {
		// conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
		// defer conn.Close()

		// ch, _ := conn.Channel()
		// defer ch.Close()

		// q, _ := ch.QueueDeclare(
		// 	"canal1", // name
		// 	false,    // durable
		// 	false,    // delete when unused
		// 	false,    // exclusive
		// 	false,    // no-wait
		// 	nil,      // arguments
		// )

		// msgs, _ := ch.Consume(
		// 	q.Name, // queue
		// 	"",     // consumer
		// 	true,   // auto-ack
		// 	false,  // exclusive
		// 	false,  // no-local
		// 	false,  // no-wait
		// 	nil,    // args
		// )

		// log.Printf("Esperando Mensajes en Canal1...")

		// for d := range msgs {
		// 	log.Printf("Received a message: %s", d.Body)
		// }
	}

	listenSecondChannel := func() {
		// conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
		// defer conn.Close()

		// ch, _ := conn.Channel()
		// defer ch.Close()

		// q, _ := ch.QueueDeclare(
		// 	"canal2", // name
		// 	false,    // durable
		// 	false,    // delete when unused
		// 	false,    // exclusive
		// 	false,    // no-wait
		// 	nil,      // arguments
		// )

		// msgs, _ := ch.Consume(
		// 	q.Name, // queue
		// 	"",     // consumer
		// 	true,   // auto-ack
		// 	false,  // exclusive
		// 	false,  // no-local
		// 	false,  // no-wait
		// 	nil,    // args
		// )

		// log.Printf("Esperando Mensajes en Canal2...")

		// for d := range msgs {
		// 	log.Printf("Received a message: %s", d.Body)
		// }
	}

	listenServer := func() {
		//creacion de enrutador pat
		router := httprouter.New()

		//Enrutadores-User
		router.GET("/", Controller.RenderIndex)

		router.GET("/lstusers.html", Controller.GetAllUser)
		router.GET("/users/addedit/:id", Controller.GetAddEditUser)
		router.POST("/users", Controller.PostPutUser)
		router.GET("/users/delete/:id", Controller.DeleteUserByID)

		//Enrutadores-Flight
		router.GET("/lstflights.html", Controller.GetAllFlight)
		router.GET("/flights/addedit/:id", Controller.GetAddEditFlight)
		router.POST("/flights", Controller.PostPutFlight)
		router.GET("/flights/delete/:id", Controller.DeleteFlightByID)

		//Enrutadores-Booking
		router.GET("/lstbookings.html", Controller.GetAllBooking)
		router.GET("/bookings/addedit/:id", Controller.GetAddEditBooking)
		router.POST("/bookings", Controller.PostPutBooking)
		router.GET("/bookings/delete/:id", Controller.DeleteBookingByID)

		//Escuchando el servidor
		log.Print("Escuchando en 127.0.0.1:9100...")
		http.ListenAndServe(":9100", router)
	}

	Parallelize(listenFirstChannel, listenSecondChannel, listenServer)
}
