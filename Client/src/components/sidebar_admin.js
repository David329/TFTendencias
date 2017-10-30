import React, { Component } from 'react'
import { Button, Dropdown, Menu, Icon } from 'semantic-ui-react'
import { withRouter, Link } from 'react-router-dom'

class SidebarMenuAdmin extends Component {

    constructor(props) {
        super(props);
        this.logout = this.logout.bind(this);
    }

    state = { activeItem: 'inicio' }

    handleItemClick = (e, { name }) => this.setState({ activeItem: name })

    logout() {
        //TODO: logica para borrar la sesión
        this.props.history.push("/");
    }

  render() {
    const { activeItem } = this.state

    return (
      <Menu inverted size='small'>
        <Menu.Item as={Link} to='/administracion/home' name='inicio' active={activeItem === 'inicio'} onClick={this.handleItemClick} className="d-none d-sm-block d-md-block d-lg-block d-xl-block">
            <Icon name='home' /> Inicio
        </Menu.Item>
        <Menu.Item as={Link} to='/administracion/vuelos' name='vuelos' active={activeItem === 'vuelos'} onClick={this.handleItemClick} className="d-none d-sm-block d-md-block d-lg-block d-xl-block">
            <Icon name='plane' /> Vuelos
        </Menu.Item>
        <Menu.Item as={Link} to='/administracion/reservas' name='reservas' active={activeItem === 'reservas'} onClick={this.handleItemClick} className="d-none d-sm-block d-md-block d-lg-block d-xl-block">
            <Icon name='ticket' /> Reservas
        </Menu.Item>
       

        <div className="d-sm-none">
            <Dropdown item text='Menú'>
                <Dropdown.Menu>
                    <Dropdown.Item as={Link} to='/administracion/home'>Inicio</Dropdown.Item>
                    <Dropdown.Item as={Link} to='/administracion/vuelos'>Vuelos</Dropdown.Item>
                    <Dropdown.Item as={Link} to='/administracion/reservas'>Reservas</Dropdown.Item>
                </Dropdown.Menu>
          </Dropdown>
        </div>

        <Menu.Menu position='right'>
          <Menu.Item>
            <Button animated inverted onClick={this.logout}>
                <Button.Content visible>Cerrar Sesión</Button.Content>
                <Button.Content hidden>
                    <Icon name='log out' />
                </Button.Content>
            </Button>
          </Menu.Item>
        </Menu.Menu>
      </Menu>
    )
  }
}

export default withRouter(SidebarMenuAdmin);