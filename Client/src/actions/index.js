import axios from 'axios'

export const VALIDAR_LOGIN = "validar_login";
export const FETCH_VUELOS = "fetch_vuelos";
export const NUEVO_VUELO = "nuevo_vuelo";
export const FETCH_PAISES = "fetch_paises";

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

    const request = axios.post(`${ROOT_URL}/flights`, values).then(callbackSuccess).catch(callbackError);

    return {
        type: NUEVO_VUELO,
        payload: request
    }
}