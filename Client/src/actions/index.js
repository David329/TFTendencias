import axios from 'axios'

export const VALIDAR_LOGIN = "validar_login";
export const FETCH_VUELOS = "fetch_vuelos";
export const NUEVO_VUELO = "nuevo_vuelo";
export const DELETE_VUELO = "delete_vuelo";
export const FETCH_PAISES = "fetch_paises";
export const FETCH_RESERVAS = "fetch_reservas";


const ROOT_URL = "http://localhost:9000"


export function validarLogin(values, callback) {

    if(typeof callback == "function") callback();

    //api de inicio de sesion
    console.log(values);

    const login = {
        status: 'ok'
    }

    return {
        type: VALIDAR_LOGIN,
        payload: login
    };
}

export function fetchPaises() {

    const request = axios.get("https://restcountries.eu/rest/v2/all");

    return {
        type: FETCH_PAISES,
        payload: request
    };
}

export function fetchVuelos() {

    const request = axios.get(`${ROOT_URL}/flights`);

    return {
        type: FETCH_VUELOS,
        payload: request
    }

}

export function nuevoVuelo(values, callbackSuccess, callbackError) {

    var data = {
        airplanemodel: values.airplanemodel,
        airplanenumber: values.airplanenumber,
        price: parseFloat(values.price),
        depart: {
            country: values.departCountry,
            city: values.departCity,
            td: values.departTD,
            ta: values.departTA
        },
        destin: {
            country: values.destinCountry,
            city: values.destinCity,
            td: values.destinTD,
            ta: values.destinTA
        },
        seats: [
            {number: 1, type: "No se"},{number: 2,type: "No se"},{number: 3,type: "No se"},
            {number: 4, type: "No se"},{number: 5,type: "No se"},{number: 6,type: "No se"},
            {number: 7, type: "No se"},{number: 8,type: "No se"},{number: 9,type: "No se"},
            {number: 10, type: "No se"},{number: 11,type: "No se"},{number: 12,type: "No se"}
        ]
    }

    const request = axios.post(`${ROOT_URL}/flights`, data).then(callbackSuccess).catch(callbackError);

    return {
        type: NUEVO_VUELO,
        payload: request
    }
}

export function eliminarVuelo(id, callbackSuccess) {
    const request = axios.delete(`${ROOT_URL}/flights/${id}`).then(callbackSuccess);

    return {
        type: NUEVO_VUELO,
        payload: request
    }
}

export function fetchReservas() {

    const request = axios.get(`${ROOT_URL}/bookings`);

    return {
        type: FETCH_RESERVAS,
        payload: request
    }
}