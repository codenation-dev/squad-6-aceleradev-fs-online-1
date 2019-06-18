import React, {Component} from 'react';
import {withRouter} from 'react-router';
import PropTypes from 'prop-types';

import alertService from './../services/alertService';
import loginService from '../services/loginService';

class Alert extends Component {
  constructor(props) {
    super(props);
    this.state = {
      alertInvalid: false,
      alertId: 0,
      alertForm: {
        id: 0,
        Arquivo: '',
        Mes: '',
        Ano:'',
        Id_pagamento:''
        },
    };
  }
 
  async componentDidMount() {
    if (this.props.match.params.id) {
      let alert;
      try {
        alert = await alertService.getAlertById(this.props.match.params.id);
      } catch (error) {
        console.log(error);
      }
      if (alert) {
        this.setState({
            alertId: this.props.match.params.id,
            alertForm: alert,
        });
      } else {
        this.setState({alertInvalid: true});
      }
    }
  }

  btnSalvarClick = async event => {
    event.preventDefault();

    let msg = '';
   
    if (!this.state.alertForm.id) {
      msg += 'Preencha o id do alerta\n';
    }

    if (!msg) {
      if (this.state.alertForm.id) {
        await alertService.putAlert(this.state.alertForm);
        this.props.history.push('/alerts');
      } else {
        await alertService.postAlert(this.state.alertForm);
        this.props.history.push('/alerts');
      }
    } else {
      alert(msg);
    }
  };

  btnCancelarClick = event => {
    event.preventDefault();
    this.props.history.push('/alerts');
  };

  handleOnChange(event) {
    if (event.target.name !== 'receiveAlert') {
      this.setState({
        ...this.state,
        alertForm: {
          ...this.state.alertForm,
          [event.target.name]: event.target.value ? event.target.value : '',
        },
      });
    } else {
      this.setState({
        ...this.state,
        alertForm: {
          ...this.state.alertForm,
          [event.target.name]: event.target.checked,
        },
      });
    }
  }

  render = () => (
    <React.Fragment>
      <script>
        console.log({JSON.stringify(loginService.userLogged(), null, 4)});
      </script>

      <div className="container">
        <div className="alert alert-danger" role="alert" hidden>
          This is a primary alert—check it out!
        </div>
        <form>
          <div className="form-row">


            <div className="form-group col-md-1">
              <label htmlFor="inputId">ID</label>
              <input
                type="text"
                className="form-control"
                placeholder="0"
                id="inputId"
                name="userId"
                value={this.state.alertId}
                disabled
              />
            </div>


            <div className="form-group col-md-3">
              <label htmlFor="inputEmail"> Data </label>
              <input
                type="text"
                className="form-control"
                placeholder="Email"
                id="inputEmail"
                name="email"
                value={this.state.alertForm.email}
                onChange={e => {
                  this.handleOnChange(e);
                }}
                required
              />
             </div>  



          
          <div className="form-group col-md-1">
              <label htmlFor="inputEmail"> Id Usuario</label>
              <input
                type="text"
                className="form-control"
                placeholder="Email"
                id="inputEmail"
                name="email"
                value={this.state.alertForm.email}
                onChange={e => {
                  this.handleOnChange(e);
                }}
                required
              />
             </div>
         
           
             <div className="form-group col-md-3">
              <label htmlFor="inputEmail">Id cliente</label>
              <input
                type="text"
                className="form-control"
                placeholder="Email"
                id="inputEmail"
                name="email"
                value={this.state.alertForm.email}
                onChange={e => {
                  this.handleOnChange(e);
                }}
                required
              />
             </div>
          

          <div className="form-group col-md-7">
            <label htmlFor="inputName">Id pagamento</label>
            <input
              type="text"
              className="form-control"
              placeholder="Nome"
              onChange={e => {
                this.handleOnChange(e);
              }}
              id="inputName"
              name="name"
              value={this.state.alertForm.name}
            />
          </div>

          </div>
          <div className="form-group">
            <button
              
              onClick={e => {
                this.btnSalvarClick(e);
              }}
              type="submit"
              className="btn btn-primary"
            >
              Salvar
            </button>
            <button
              onClick={e => {
                this.btnCancelarClick(e);
              }}
              type="submit"
              className="btn btn-secondary"
            >
              Cancelar
            </button>
          </div>
        </form>
      </div>
    </React.Fragment>
  );

  static propTypes = {
    location: PropTypes.object.isRequired,
    history: PropTypes.object.isRequired,
  };
}

export default withRouter(Alert);
