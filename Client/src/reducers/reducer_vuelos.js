import { FETCH_VUELOS } from "../actions"

export default function(state = null, action) {

    switch (action.type) {
        case FETCH_VUELOS:
            return action.payload.data;
            break;
    
        default:
            return state;
            break;
    }
}