import { FETCH_PAISES } from "../actions"

export default function(state = [], action) {

    switch (action.type) {
        case FETCH_PAISES:
            const paisesEsp = action.payload.data.map(x=> x.translations.es);
            return paisesEsp.filter((pais) => { return pais != null });
            break;
    
        default:
            return state;
            break;
    }
}