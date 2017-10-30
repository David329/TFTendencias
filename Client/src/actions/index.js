import axios from 'axios'

export const VALIDAR_LOGIN = "validar_login";
export const FETCH_VUELOS = "fetch_vuelos";

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

export function fetchVuelos() {

    const request = axios.get(`${ROOT_URL}/flights`);

    return {
        type: FETCH_VUELOS,
        payload: request
    }

}