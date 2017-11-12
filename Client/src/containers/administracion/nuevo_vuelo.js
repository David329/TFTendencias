import React, { Component } from 'react'
import { Field, reduxForm } from 'redux-form'
import { Link } from 'react-router-dom'
import { connect } from 'react-redux'
import { nuevoVuelo } from '../../actions'
import { fetchPaises } from '../../actions';
import SelectField from 'material-ui/SelectField';
import MenuItem from 'material-ui/MenuItem';
import {Card, CardActions, CardHeader} from 'material-ui/Card';
import TextField from 'material-ui/TextField';
import { Button, Icon } from 'semantic-ui-react'
import { Divider } from 'semantic-ui-react'
import TimePicker from 'material-ui/TimePicker';
import { ToastContainer, toast } from 'react-toastify';

class VueloNuevo extends Component {

    constructor(props) {
        super(props);
        this.state = {
            paisSeleccionadoOrigen: null,
            paisSeleccionadoDestino: null
        }
    }

    componentDidMount() {
        this.props.fetchPaises();
    }

    notifySuccess() {
        toast.success("Vuelo creado exitosamente.", {
          position: toast.POSITION.TOP_CENTER,
          autoClose: 3000,
          onClose: (childrenProps) =>  this.props.history.push("/administracion/vuelos")
        });
    }

    notifyError() {
        toast.error("Hubo un problema en la operación.", {
            position: toast.POSITION.TOP_CENTER,
            autoClose: 3000
          });
    }


    renderPaises() {

        return this.props.paises.map((pais) => {

            //if(!pais) return;

            return(
                <MenuItem value={pais} key={pais} primaryText={pais} />
            );
        });
    }

    handleChangePaises = (event, index, value) => {
        this.setState({paisSeleccionadoOrigen: value});
    };

    renderTextField(field) {

        const { meta: { touched, error } } = field;

        return(
            <TextField {...field.input} hintText={field.hintText} floatingLabelText={field.label} fullWidth={true} type={field.type} errorText={touched ? error : ''} />
        );
    }

    renderDropDown(field) {

        const { meta: { touched, error } } = field;

        return (
            <SelectField 
                {...field.input}
                maxHeight ={500}
                hintText = {field.hintText}
                onChange={(event, index, value) => field.input.onChange(value)}
                errorText={touched ? error : ''}
                style={{'marginTop': '20px'}}
                fullWidth={true}
            >
                {field.options}
            </SelectField>
        );
    }

    onSubmit(values) {
        
        this.props.nuevoVuelo(values, () => {
            this.notifySuccess();
        }, () => {
            this.notifyError();
        })
    }


    render() {
        
        const { handleSubmit } = this.props;

        return(
            <div className="row">
                <div className="col-sm-12">
                <form onSubmit={handleSubmit(this.onSubmit.bind(this))}>
                    <Card>
                        <CardHeader 
                            title="Nuevo Vuelo"
                            titleStyle={{'fontSize':'20px', 'fontWeight':'bold'}}
                        />
                        <CardActions>
                            <b>General</b>
                            <Divider className="mt-0"></Divider>
                            <div className="row mb-5 px-4">
                                <div className="col-sm-4">
                                    <Field name="airplanemodel" hintText="Ingrese Modelo" label="Modelo" type="text" component={this.renderTextField}/>
                                </div>
                                <div className="col-sm-4">
                                    <Field name="airplanenumber" label="Número" type="number" component={this.renderTextField}/>
                                </div>
                                <div className="col-sm-4">
                                    <Field name="price" label="Precio" type="text" component={this.renderTextField}/>
                                </div>
                            </div>
                            <b>Origen</b>
                            <Divider className="mt-0"></Divider>
                            <div className="row mb-5 px-4">
                                <div className="row col-sm-12">
                                    <div className="col-sm-4">
                                        <Field name="departCountry" hintText="País" component={this.renderDropDown} options={this.renderPaises()}/>
                                    </div>
                                    <div className="col-sm-4">
                                        <Field name="departCity" label="Ciudad" hintText="Ingrese ciudad" type="text" component={this.renderTextField}/>
                                    </div>       
                                    <div className="col-sm-2">
                                    <Field name="departTD" hintText="Ingrese TD" label="TD" type="text" component={this.renderTextField}/>
                                    </div>        
                                    <div className="col-sm-2">
                                    <Field name="departTA" hintText="Ingrese TA" label="TA" type="text" component={this.renderTextField}/>
                                    </div>                           
                                </div>
                            </div>
                            <b>Destino</b>
                            <Divider className="mt-0"></Divider>
                            <div className="row px-4">
                                <div className="row col-sm-12">
                                    <div className="col-sm-4">
                                        <Field name="destinCountry" hintText="País" component={this.renderDropDown} options={this.renderPaises()}/>
                                    </div>
                                    <div className="col-sm-4">
                                        <Field name="destinCity" label="Ciudad" hintText="Ingrese ciudad" type="text" component={this.renderTextField}/>
                                    </div>       
                                    <div className="col-sm-2">
                                        <Field name="destinTD" hintText="Ingrese TD" label="TD" type="text" component={this.renderTextField}/>
                                    </div>        
                                    <div className="col-sm-2">
                                        <Field name="destinTA" hintText="Ingrese TA" label="TA" type="text" component={this.renderTextField}/>
                                    </div>                           
                                </div>
                            </div>
                        </CardActions>
                    </Card>

                    <Card className="mt-4">
                        <CardActions>
                            <Button animated primary type="submit">
                                <Button.Content visible>Guardar</Button.Content>
                                <Button.Content hidden>
                                    <Icon name='save' />
                                </Button.Content>
                            </Button>
                            <Button animated secondary as={Link} to="/administracion/vuelos">
                                <Button.Content visible>Volver</Button.Content>
                                <Button.Content hidden>
                                    <Icon name='arrow circle left' />
                                </Button.Content>
                            </Button>
                        </CardActions>
                    </Card>
                    <ToastContainer/>
                </form>
                </div>
            </div>
        );
    }
}

function validate(values) {
    const errors = {}
    return errors;
}

function mapStateToProps(state) {
    return {paises: state.paises}
}

export default reduxForm({
    form: 'NuevoVueloForm',
    validate
})(
    connect(mapStateToProps, { fetchPaises, nuevoVuelo })(VueloNuevo)
);