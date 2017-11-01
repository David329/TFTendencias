import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux'
import { createStore, applyMiddleware } from 'redux'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import ReduxPromise from 'redux-promise'

import './lib/styles/index.css';
import 'semantic-ui-css/semantic.min.css';
import 'react-toastify/dist/ReactToastify.min.css';

import registerServiceWorker from './registerServiceWorker';

import reducers from './reducers'
import Login from './containers/login/login'
import AdministracionLayout from './containers/administracion/layout'
import AdministracionHome from './containers/administracion/home'
import AdministracionVuelos from './containers/administracion/vuelos'
import VueloNuevo from './containers/administracion/nuevo_vuelo'

const createStoreWithMiddleware = applyMiddleware(ReduxPromise)(createStore);

ReactDOM.render(
    <Provider store={createStoreWithMiddleware(reducers)}>
        <BrowserRouter>
        <div>
            <Switch>
                <AdministracionLayout exact path="/administracion/home" component={AdministracionHome} />
                <AdministracionLayout exact path="/administracion/vuelos" component={AdministracionVuelos} />
                <AdministracionLayout exact path="/administracion/vuelos/nuevo" component={VueloNuevo} />
                <Route exact path="/" component={Login} />
            </Switch>
        </div>
        </BrowserRouter>
    </Provider>, 
    document.getElementById('root')
);

registerServiceWorker();
