import { FETCH_RESERVAS } from "../actions"

export default function(state = null, action) {

    switch (action.type) {
        case FETCH_RESERVAS:
            return action.payload.data;
            break;
    
        default:
            return state;
            break;
    }
}