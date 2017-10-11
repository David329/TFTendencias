package controller

//Restful - Flight
import (
	"encoding/json"
	"log"
	"net/http"

	Models "../Model"

	"github.com/julienschmidt/httprouter"
)

//Get-Post-Put-Delete

//GetAllFlight Envia todos los vuelos, formato->JSON
func GetAllFlight(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	req, err := http.NewRequest("GET", "http://localhost:8000/flights", nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var obj []Models.Flight

	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		log.Println(err)
	}

	//	wr.Header().Set("Content-Type", "text/html")

	//	fmt.Print(wr, string(responseData))
}

//PostFlight Inserta un nuevo vuelo
func PostFlight(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {
}

//PutFlightByID Actualiza un Documento Flight
func PutFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
}

//DeleteFlightByID Elimina un usuario por ID, formato->JSON
func DeleteFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {
}
