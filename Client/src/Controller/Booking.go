package controller

//Restful - Booking
import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	Models "../Model"
	"github.com/julienschmidt/httprouter"
)

//Get-Post-Put-Delete

//GetAllBooking Envia todos las reservar, formato->HTML
func GetAllBooking(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var obj []Models.Booking
	response, err := http.Get("http://localhost:8000/bookings")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &obj)

	testTemplate, _ := template.ParseFiles("./View/bookings.gohtml")

	wr.Header().Set("Content-Type", "text/html")

	err = testTemplate.Execute(wr, obj)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

// //PostBooking Inserta un nuevo vuelo
// func PostBooking(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

// }

// //PutBookingByID Actualiza un Documento Booking
// func PutBookingByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) { //pensar si es correcto...

// }

// //DeleteBookingByID Elimina un usuario por ID, formato->JSON
// func DeleteBookingByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

// }
