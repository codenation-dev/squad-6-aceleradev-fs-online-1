import React, {Component} from 'react';
import {withRouter} from 'react-router';
import PropTypes from 'prop-types';

import customersService from './../services/customerService';
import loginService from '../services/loginService';

class Alert extends Component {
  constructor(props) {
    super(props);
    this.state = {
      customerInvalid: false,
      customerId: 0,
      customerForm: {
        id: 0,
        email: '',
        name: '',
        },
    };
  }
 
  async componentDidMount() {
    if (this.props.match.params.id) {
      let customer;
      try {
        customer = await customersService.getCustomerById(this.props.match.params.id);
      } catch (error) {
        console.log(error);
      }
      if (customer) {
        this.setState({
            customerId: this.props.match.params.id,
            customerForm: customer,
        });
      } else {
        this.setState({customerInvalid: true});
      }
    }
  }

  btnSalvarClick = async event => {
    event.preventDefault();

    let msg = '';
   
    if (!this.state.customerForm.name) {
      msg += 'Preencha o campo Nome\n';
    }

    if (!msg) {
      if (this.state.customerForm.id) {
        await customersService.putCustomer(this.state.customerForm);
        this.props.history.push('/customers');
      } else {
        await customersService.postCustomer(this.state.customerForm);
        this.props.history.push('/customers');
      }
    } else {
      alert(msg);
    }
  };

  btnCancelarClick = event => {
    event.preventDefault();
    this.props.history.push('/customers');
  };

  handleOnChange(event) {
    if (event.target.name !== 'receiveAlert') {
      this.setState({
        ...this.state,
        customerForm: {
          ...this.state.customerForm,
          [event.target.name]: event.target.value ? event.target.value : '',
        },
      });
    } else {
      this.setState({
        ...this.state,
        customerForm: {
          ...this.state.customerForm,
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
            <div className="form-group col-md-2">
              <label htmlFor="inputId">ID</label>
              <input
                type="text"
                className="form-control"
                placeholder="0"
                id="inputId"
                name="userId"
                value={this.state.customerId}
                disabled
              />
            </div>
            <div className="form-group col-md-7">
              <label htmlFor="inputEmail">Email</label>
              <input
                type="text"
                className="form-control"
                placeholder="Email"
                id="inputEmail"
                name="email"
                value={this.state.customerForm.email}
                onChange={e => {
                  this.handleOnChange(e);
                }}
                required
              />
            
            </div>
          </div>
          <div className="form-group">
            <label htmlFor="inputName">Nome</label>
            <input
              type="text"
              className="form-control"
              placeholder="Nome"
              onChange={e => {
                this.handleOnChange(e);
              }}
              id="inputName"
              name="name"
              value={this.state.customerForm.name}
            />
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
