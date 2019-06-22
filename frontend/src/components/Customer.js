import React, {Component} from 'react';
import {withRouter} from 'react-router';
import PropTypes from 'prop-types';

import customersService from './../services/customerService';
import paymentsService from './../services/paymentsService';

class Customer extends Component {
  constructor(props) {
    super(props);
    this.state = {
      customerInvalid: false,
      customerId: 0,
      customerForm: {
        id: 0,
        name: '',
      },
      listPaymentsCustomer: [],
    };
  }

  async componentDidMount() {
    if (this.props.match.params.id) {
      let listPaymentsCustomer = [];
      let customer;
      try {
        customer = await customersService.getCustomerById(
          this.props.match.params.id
        );
        if (customer.id) {
          listPaymentsCustomer = await paymentsService.getPayments(customer.id);
        }
      } catch (error) {
        console.log(error);
      }
      if (customer) {
        this.setState({
          customerId: this.props.match.params.id,
          customerForm: customer,
          listPaymentsCustomer: listPaymentsCustomer,
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
      <div className="container">
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
        <br />
        <div className="row">
          {this.state.listPaymentsCustomer ? (
            <React.Fragment>
              <h4>Histórico de Pagamentos pelo Governo de SP</h4>
              <table className="table table-hover">
                <thead>
                  <tr>
                    <th scope="col" hidden>
                      #
                    </th>
                    <th scope="col">Ano-Mês</th>
                    <th scope="col">Departamento</th>
                    <th scope="col">Cargo</th>
                    <th scope="col">Salário</th>
                  </tr>
                </thead>
                <tbody>
                  {this.state.listPaymentsCustomer
                    ? this.state.listPaymentsCustomer.map((payment, index) =>
                        payment.EmployeePayments.map((item, index2) => (
                          <tr key={item.id}>
                            <th scope="row" hidden>
                              {item.id}
                            </th>
                            <th>
                              {String(payment.year + '-' + payment.month)}
                            </th>
                            <th>{item.occupation}</th>
                            <th>{item.department}</th>
                            <th>
                              {item.salary.toLocaleString('pt-br', {
                                style: 'currency',
                                currency: 'BRL',
                              })}
                            </th>
                          </tr>
                        ))
                      )
                    : ''}
                </tbody>
              </table>
            </React.Fragment>
          ) : (
            ''
          )}
        </div>
      </div>
    </React.Fragment>
  );

  static propTypes = {
    location: PropTypes.object.isRequired,
    history: PropTypes.object.isRequired,
  };
}

export default withRouter(Customer);
