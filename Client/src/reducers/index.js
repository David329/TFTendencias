import { combineReducers } from 'redux';
import { reducer as formReducer } from 'redux-form'
import VuelosReducer from './reducer_vuelos'
import PaisesReducer from './reducer_paises'

const rootReducer = combineReducers({
  form: formReducer,
  vuelos: VuelosReducer,
  paises: PaisesReducer
});

export default rootReducer;