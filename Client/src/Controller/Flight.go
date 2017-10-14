package controller

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

//GetAllFlight Envia todos los vuelos, formato->HTML
func GetAllFlight(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var obj []Models.Flight
	response, err := http.Get("http://localhost:8000/flights")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &obj)

	tmpl, _ := template.ParseFiles(
		"./View/templates/header.gohtml",
		"./View/Flights/lstflights.gohtml",
		"./View/templates/footer.gohtml",
	)
	wr.Header().Set("Content-Type", "text/html")

	err = tmpl.ExecuteTemplate(wr, "lstflights", obj)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}
}

// //PostFlight Inserta un nuevo vuelo
// func PostFlight(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {
// }

// //PutFlightByID Actualiza un Documento Flight
// func PutFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
// }

//DeleteFlightByID Elimina un usuario por ID, formato->JSON
func DeleteFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	response, _ := http.NewRequest("DELETE", "http://localhost:8000/flights/"+ps.ByName("id"), nil)

	new(http.Client).Do(response)

	GetAllFlight(wr, req, ps)
}
