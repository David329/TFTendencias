import React, { Component } from 'react';
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'

import {Card, CardActions, CardHeader} from 'material-ui/Card';
import { Button, Icon } from 'semantic-ui-react';
import MenuItem from 'material-ui/MenuItem';
import IconButton from 'material-ui/IconButton';
import IconMenu from 'material-ui/IconMenu';
import { fetchReservas } from '../../actions';

import {
    Table,
    TableBody,
    TableHeader,
    TableHeaderColumn,
    TableRow,
    TableRowColumn,
  } from 'material-ui/Table';


class AdministracionReservas extends Component {

    componentDidMount() {
        this.props.fetchReservas();
    }

    renderReservas() {

        if(!this.props.reservas) return <TableRow><TableRowColumn><div><h4>Cargando...</h4></div></TableRowColumn></TableRow>;

        return this.props.reservas.map((reserva) => {
            return (
                <TableRow key={reserva.ID}>
                    <TableRowColumn>{reserva.UserID}</TableRowColumn>
                    <TableRowColumn>{reserva.FlightID}</TableRowColumn>
                    <TableRowColumn>{reserva.PersonalSeat.Number}</TableRowColumn>
                    <TableRowColumn>{reserva.PersonalSeat.Type}</TableRowColumn>
                </TableRow>
            );
        });
    }

    render() {
        return(
            <div>
                <Card>
                    <CardHeader 
                        title="Reservas"
                        subtitle="Lista de los reservas realizadas"
                    />
                    <CardActions>
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
                                    <TableHeaderColumn>Usuario</TableHeaderColumn>
                                    <TableHeaderColumn>Vuelo</TableHeaderColumn>
                                    <TableHeaderColumn>Sitio</TableHeaderColumn>
                                    <TableHeaderColumn>Tipo</TableHeaderColumn>
                                </TableRow>
                            </TableHeader>
                            <TableBody
                                displayRowCheckbox={false}
                            >
                                {this.renderReservas()}                                
                            </TableBody>
                        </Table>
                    </CardActions>
                </Card>
           </div>
        );
    }
}

function mapStateToProps(state) {
    return {reservas: state.reservas}
}

export default connect(mapStateToProps, { fetchReservas })(AdministracionReservas);