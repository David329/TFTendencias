import React, { Component } from 'react';
import { connect } from 'react-redux'

import {Card, CardActions, CardHeader} from 'material-ui/Card';
import { Button, Icon } from 'semantic-ui-react';
import MenuItem from 'material-ui/MenuItem';
import IconButton from 'material-ui/IconButton';
import IconMenu from 'material-ui/IconMenu';
import MenuIcon from 'material-ui/svg-icons/navigation/menu';
import { fetchVuelos } from '../../actions';

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

    renderVuelos() {

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
                            <MenuItem value="1" primaryText="Detalle" />
                            <MenuItem value="2" primaryText="Editar" />
                            <MenuItem value="3" primaryText="Eliminar" />
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
                        <Button primary animated className="float-right" size='small'>
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
           </div>
        );
    }
}

function mapStateToProps(state) {
    return {vuelos: state.vuelos}
}

export default connect(mapStateToProps, { fetchVuelos })(AdministracionVuelos);