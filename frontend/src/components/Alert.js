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
        filename: '',
        month: '',
        year:'',
        EmployeePayments:''
        },
    };
  }
 
  async componentDidMount() {
    if (this.props.match.params.id) {
      let alert;
      try {
        alert = await alertService.getAlertById(this.props.match.params.id);
        console.log(alert,"aqui nos alertas")
      } catch (error) {
        console.log(error);
      }
      if (alert) {
        this.setState({
            alertId: this.props.match.params.id,
            alertForm:  alert[0]
        });
      } else {
        this.setState({alertInvalid: true});
      }
    }
  }

  btnSalvarClick = async event => {
    event.preventDefault();

    let msg = '';
    if (!this.state.alertForm.filename) {
      msg += 'Preencha o campo Arquivo\n';
    }
    if (!this.state.alertForm.Mes) {
      msg += 'Preencha o campo Mês\n';
    }

    if (!this.state.alertForm.Ano) {
      msg += 'Preencha o campo Ano\n';
    }
    if (!this.state.alertForm.Id_pagamento) {
      msg += 'Preencha o campo detalhes dos salarios \n';
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
  
      this.setState({
        ...this.state,
        alertForm: {
          ...this.state.alertForm,
          [event.target.name]: event.target.value ? event.target.value : '',
        },
      });
    
    }
  

  render = () => (
    <React.Fragment>
      <script>
       {/*  console.log({JSON.stringify(loginService.userLogged(), null, 4)}); */}
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
                name="alertId"
                value={this.state.alertId}
                disabled
              />
            </div>


            <div className="form-group col-md-8">
              <label htmlFor="inputEmail"> Arquivo </label>
              <input
                type="text"
                className="form-control"
                placeholder="Arquivo"
                id="inputarquivo"
                name="email"
                value={this.state.alertForm.filename}
               
                onChange={e => {
                  this.handleOnChange(e);
                }}
                disabled
              />
              
             </div>  



          
          <div className="form-group col-md-1">
              <label htmlFor="inputEmail"> Mês</label>
              <input
                type="text"
                className="form-control"
                placeholder="Mês"
                id="inputEmail"
                name="email"
                value={this.state.alertForm.month}
                onChange={e => {
                  this.handleOnChange(e);
                }}
                disabled
              />
             </div>
         
           
             <div className="form-group col-md-1">
              <label htmlFor="inputEmail">Ano</label>
              <input
                type="text"
                className="form-control"
                placeholder="Ano"
                id="inputEmail"
                name="email"
                value={this.state.alertForm.year}
                onChange={e => {
                  this.handleOnChange(e);
                }}
                disabled
              />
             </div>
          

          <div className="form-group col-md-7">
            <label htmlFor="inputName">Salário Cód. Nº</label>
            <input
              type="text"
              className="form-control"
              placeholder="Salário Cód."
              onChange={e => {
                this.handleOnChange(e);
              }}
              id="inputName"
              name="name"
              value={this.state.alertForm.EmployeePayments}
              disabled
            />
          </div>

          </div>
          <div className="form-group">
           {/*  <button
              
              onClick={e => {
                this.btnSalvarClick(e);
              }}
              type="submit"
              className="btn btn-primary"
            >
              Salvar
            </button> */}
            <button
              onClick={e => {
                this.btnCancelarClick(e);
              }}
              type="submit"
              className="btn btn-secondary"
            >
              Voltar
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
