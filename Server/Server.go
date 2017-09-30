////DOCUMENTATION
//https://godoc.org/gopkg.in/mgo.v2
//http://www.jancarloviray.com/blog/go-mongodb-simple-example/
//PREGUNTAR LAS CREDENCIALES DE USUARIO DE BASE DE DATOS
package main

//MEJORAR:
//  1. Pasar los parametros como direccion de memoria, para eliminar la sesion y variables
//  2. Por ahora estamos enviando desde ./Test/SendDotNet/Send#, esto es solo de prueba, xq debe ser enviado desde el cliente.
import (
	"log"
	"net/http"
	"sync"

	Routes "./Routes"

	"github.com/gorilla/pat"
	"github.com/streadway/amqp"
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
		conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
		defer conn.Close()

		ch, _ := conn.Channel()
		defer ch.Close()

		q, _ := ch.QueueDeclare(
			"canal1", // name
			false,    // durable
			false,    // delete when unused
			false,    // exclusive
			false,    // no-wait
			nil,      // arguments
		)

		msgs, _ := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}

	listenSecondChannel := func() {
		conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
		defer conn.Close()

		ch, _ := conn.Channel()
		defer ch.Close()

		q, _ := ch.QueueDeclare(
			"canal2", // name
			false,    // durable
			false,    // delete when unused
			false,    // exclusive
			false,    // no-wait
			nil,      // arguments
		)

		msgs, _ := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}

	listenServer := func() {
		//creacion de enrutador pat
		router := pat.New()

		//llamada de metodos a enrutar
		router.Get("/users", Routes.GetAllUser)
		router.Delete("/users/{id}", Routes.DeleteUserByID)
		router.Post("/flights", Routes.PostFlight)
		router.Put("/flights/{id}", Routes.PutFlightByID)

		//activar entutador, probar a;adirlo en otro archivo...
		http.Handle("/", router)

		//escuchando el servidor
		log.Print("Escuchando en 127.0.0.1:8000...")
		log.Fatal(http.ListenAndServe(":8000", nil))
	}

	Parallelize(listenFirstChannel, listenSecondChannel, listenServer)
}
