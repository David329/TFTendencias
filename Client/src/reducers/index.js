import { combineReducers } from 'redux';
import { reducer as formReducer } from 'redux-form'
import VuelosReducer from './reducer_vuelos'

const rootReducer = combineReducers({
  form: formReducer,
  vuelos: VuelosReducer
});

export default rootReducer;