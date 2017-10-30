import React, { Component } from 'react'
import { Route } from 'react-router-dom'
import AppBar from 'material-ui/AppBar';

import {MuiThemeProvider} from 'material-ui/styles'
import SidebarMenuAdmin from '../../components/sidebar_admin'


const AdministracionLayout = ({component: Component, ...rest}) => {

    return (
      <Route {...rest} render={matchProps => (
        <div>
            <SidebarMenuAdmin />
            <MuiThemeProvider>
            <div className="container mt-5">
                <Component {...matchProps} />
            </div>
            </MuiThemeProvider>
        </div>
      )} />
    )
};

export default AdministracionLayout