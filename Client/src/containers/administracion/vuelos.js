import React, { Component } from 'react';
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'

import {Card, CardActions, CardHeader} from 'material-ui/Card';
import { Button, Icon } from 'semantic-ui-react';
import MenuItem from 'material-ui/MenuItem';
import IconButton from 'material-ui/IconButton';
import IconMenu from 'material-ui/IconMenu';
import MenuIcon from 'material-ui/svg-icons/navigation/menu';
import { fetchVuelos, eliminarVuelo } from '../../actions';
import { ToastContainer, toast } from 'react-toastify';

import {
    Table,
    TableBody,
    TableHeader,
    TableHeaderColumn,
    TableRow,
    TableRowColumn,
  } from 'material-ui/Table';

class AdministracionVuelos extends Component {

    componentDidMount() {
        this.props.fetchVuelos();
    }

    notifySuccess() {
        toast.success("Vuelo eliminado exitosamente.", {
          position: toast.POSITION.TOP_CENTER,
          autoClose: 3000
        });
    }

    renderVuelos() {

        if(!this.props.vuelos) return <TableRow><TableRowColumn><div><h4>Cargando...</h4></div></TableRowColumn></TableRow>;

        return this.props.vuelos.map((vuelo) => {
            return (
                <TableRow key={vuelo.ID}>
                    <TableRowColumn>{vuelo.AirplaneNumber}</TableRowColumn>
                    <TableRowColumn>{vuelo.Price}</TableRowColumn>
                    <TableRowColumn>{vuelo.Depart.City}</TableRowColumn>
                    <TableRowColumn>{vuelo.Destin.City}</TableRowColumn>
                    <TableRowColumn>
                        <IconMenu
                            iconButtonElement={<IconButton><MenuIcon /></IconButton>}
                            useLayerForClickAway={true}
                            anchorOrigin={{horizontal: 'left', vertical: 'bottom'}}
                            targetOrigin={{horizontal: 'left', vertical: 'top'}}
                        >
                            <MenuItem value="1">
                                <div onClick={() => {console.log("holo")}}>Detalle</div>
                            </MenuItem>
                            <MenuItem value="2" >
                                <div onClick={() => {console.log("holo")}}>Editar</div>
                            </MenuItem>
                            <MenuItem value="3">
                                <div onClick={() => this.props.eliminarVuelo(vuelo.ID, () => { this.props.fetchVuelos(); this.notifySuccess(); })}>Eliminar</div>
                            </MenuItem>
                        </IconMenu>
                    </TableRowColumn>
                </TableRow>
            );
        });
    }

    render() {
        return(
           <div>
                <Card>
                    <CardHeader 
                        title="Vuelos"
                        subtitle="Lista de los vuelos de la aerolínea"
                    />
                    <CardActions>
                    <div className="mb-4">
                        <Button as={Link} to="/administracion/vuelos/nuevo" primary animated className="float-right" size='small'>
                            <Button.Content visible>Nuevo</Button.Content>
                            <Button.Content hidden>
                                <Icon name='add' />
                            </Button.Content>
                        </Button>
                    </div>
                    </CardActions>
                    <CardActions>
                        <Table 
                            selectable={false}
                            multiSelectable={false}
                            fixedFooter={false}
                            fixedHeader={true}
                        >
                            <TableHeader 
                                adjustForCheckbox={false}
                                displaySelectAll={false}
                                enableSelectAll={false}
                            >
                                <TableRow>
                                    <TableHeaderColumn>Avión</TableHeaderColumn>
                                    <TableHeaderColumn>Precio</TableHeaderColumn>
                                    <TableHeaderColumn>Origen</TableHeaderColumn>
                                    <TableHeaderColumn>Destino</TableHeaderColumn>
                                    <TableHeaderColumn>Opciones</TableHeaderColumn>
                                </TableRow>
                            </TableHeader>
                            <TableBody
                                displayRowCheckbox={false}
                            >
                                {this.renderVuelos()}                                
                            </TableBody>
                        </Table>
                    </CardActions>
                </Card>
                <ToastContainer/>
           </div>
        );
    }
}

function mapStateToProps(state) {
    return {vuelos: state.vuelos}
}

export default connect(mapStateToProps, { fetchVuelos, eliminarVuelo })(AdministracionVuelos);