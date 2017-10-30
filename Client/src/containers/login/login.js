import React, { Component } from 'react'
import { Field, reduxForm } from 'redux-form'
import { connect } from 'react-redux'
import { validarLogin } from '../../actions'

import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import {MuiThemeProvider} from 'material-ui/styles'
import {Card, CardActions, CardHeader} from 'material-ui/Card';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
import { ToastContainer, toast } from 'react-toastify';


import '../../lib/styles/login.css'

class Login extends Component {

    notifyCredencialesIncorrectas = () => {
      toast.error("¡Credenciales incorrectas!", {
        position: toast.POSITION.TOP_CENTER,
        autoClose: 3000
      });
    };

    renderField(field) {

        const { meta: { touched, error } } = field;

        return(
            <div>
                <TextField {...field.input} hintText={field.hintText} floatingLabelText={field.label} fullWidth={true} type={field.type} errorText={touched ? error : ''} />
            </div>
        );
    }

    onSubmit(values) {
        this.props.validarLogin(values, () => {
            this.props.history.push("/administracion/home");
        });

        this.props.history.push("/administracion/home");
    }

    render() {

        const { handleSubmit } = this.props;

        return(
            <div className="loginContainer mx-auto">
                <form onSubmit={handleSubmit(this.onSubmit.bind(this))}>
                    <MuiThemeProvider muiTheme={getMuiTheme(darkBaseTheme)}>
                        <Card>
                            <CardHeader 
                                title="Inicio de Sesión"
                                subtitle="Ingrese sus credenciales"
                            />
                            <CardActions>
                                <Field name="usuario" hintText="Ingrese Usuario" label="Usuario" type="text" component={this.renderField}/>
                                <Field name="password" hintText="Ingrese Password" label="Password" type="password" component={this.renderField}/>
                                <RaisedButton type="submit" className="mt-2" label="Login" primary={true} />
                            </CardActions>
                        </Card>
                    </MuiThemeProvider>
                </form>
                <ToastContainer/>
            </div>
        );
    }
}

function validate(values) {
    const errors = {}

    if(!values.usuario) {
        errors.usuario = "El usuario es obligatorio."
    }

    if(!values.password) {
        errors.password = "El password es obligatorio."
    }

    return errors;
}

export default reduxForm({
    form: 'LoginForm',
    validate
})(
    connect(null, { validarLogin })(Login)
);