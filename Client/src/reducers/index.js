import { combineReducers } from 'redux';
import { reducer as formReducer } from 'redux-form'
import VuelosReducer from './reducer_vuelos'
import PaisesReducer from './reducer_paises'
import ReservasReducer from './reducer_reservas'

const rootReducer = combineReducers({
  form: formReducer,
  vuelos: VuelosReducer,
  reservas: ReservasReducer
});

export default rootReducer;