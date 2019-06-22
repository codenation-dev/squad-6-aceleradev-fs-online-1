import React, {Component} from 'react';
import {withRouter} from 'react-router';
import PropTypes from 'prop-types';

import dayjs from 'dayjs';
import qs from 'query-string';

import alertService from '../services/alertService';
import usersService from '../services/usersService';
import paymentsService from '../services/paymentsService';

class Alerts extends Component {
  constructor(props) {
    super(props);
    this.state = {
      listAlerts: [],
      listUsers: [],
      listPayments: [],
      filters: {onlyCustomers: 1, userId: 0, paymentId: 0},
    };
  }
  async componentDidMount() {
    const paramsQuery =
      this.props.location && this.props.location.search
        ? qs.parse(this.props.location.search)
        : {onlyCustomers: 1, userId: 0, paymentId: 0};

    this.setState({
      listAlerts: await alertService.getAlerts(qs.stringify(paramsQuery)),
      listUsers: await usersService.getUsers(),
      listPayments: await paymentsService.getPayments(),

      filters: {
        onlyCustomers: paramsQuery.onlyCustomers
          ? paramsQuery.onlyCustomers
          : 0,
        userId: paramsQuery.userId ? paramsQuery.userId : 0,
        paymentId: paramsQuery.paymentId ? paramsQuery.paymentId : 0,
      },
    });
  }

  async handleChange(event) {
    if (event.target.id === 'inputApenasClientes') {
      this.setState(
        {
          filters: {...this.state.filters, onlyCustomers: event.target.value},
        },
        async () => {
          const params = qs.stringify(this.state.filters);
          this.props.history.push('/alerts?' + params);
          this.setState({
            listAlerts: await alertService.getAlerts(params),
          });
        }
      );
    } else if (event.target.id === 'inputApenasClientes') {
      this.setState(
        {
          filters: {...this.state.filters, onlyCustomers: event.target.value},
        },
        async () => {
          const params = qs.stringify(this.state.filters);
          this.props.history.push('/alerts?' + params);
          this.setState({
            listAlerts: await alertService.getAlerts(params),
          });
        }
      );
    } else if (event.target.id === 'inputUser') {
      this.setState(
        {
          filters: {...this.state.filters, userId: event.target.value},
        },
        async () => {
          const params = qs.stringify(this.state.filters);
          this.props.history.push('/alerts?' + params);
          this.setState({
            listAlerts: await alertService.getAlerts(params),
          });
        }
      );
    }
  }

  render = () => (
    <React.Fragment>
      <div className="container">
        <div className="row">
          <h5>Filtros</h5>
        </div>
        <div className="row">
          <div className="input-group mb-3">
            <div className="input-group-prepend">
              <label className="input-group-text" htmlFor="inputApenasClientes">
                Apenas Clientes?
              </label>
            </div>
            <select
              className="custom-select"
              id="inputApenasClientes"
              onChange={e => this.handleChange(e)}
              value={this.state.filters.onlyCustomers}
            >
              <option value="0">Não</option>
              <option value="1">Sim</option>
            </select>
            <div className="input-group-prepend">
              <label className="input-group-text" htmlFor="inputUser">
                Usuário
              </label>
            </div>
            <select
              className="custom-select"
              id="inputUser"
              onChange={e => this.handleChange(e)}
              value={this.state.filters.userId}
            >
              <option value="0">Mostrar Tudo</option>
              {this.state.listUsers
                ? this.state.listUsers.map((user, index) => (
                    <option key={user.id} value={user.id}>
                      {user.name}
                    </option>
                  ))
                : ''}
            </select>
            <div className="input-group-prepend">
              <label
                className="input-group-text"
                htmlFor="inputPayment"
                onChange={e => this.handleChange(e)}
                value={this.state.filters.paymentId}
              >
                Pagamento
              </label>
            </div>
            <select className="custom-select" id="inputPayment">
              <option value="">Mostrar Tudo</option>
              {this.state.listPayments
                ? this.state.listPayments.map((payment, index) => (
                    <option key={payment.id} value={payment.id}>
                      {String(payment.year + '-' + payment.month)}
                    </option>
                  ))
                : ''}
            </select>
          </div>
        </div>

        <br />
        <div className="row">
          {this.state.listAlerts ? (
            <table className="table table-sm table-hover">
              <thead>
                <tr>
                  <th scope="col">#</th>
                  <th scope="col">Data\Hora</th>
                  <th scope="col">Usuário Alerta</th>
                  <th scope="col">Usuário Email</th>
                  <th scope="col">Cliente?</th>
                  <th scope="col">Pagamento</th>
                  <th scope="col">Funcionário Gov. SP</th>
                  <th scope="col">Salário</th>
                </tr>
              </thead>

              <tbody>
                {this.state.listAlerts
                  ? this.state.listAlerts.map((item, index) => (
                      <tr key={item.id}>
                        <th scope="row">{item.id}</th>
                        <td>
                          {dayjs(item.date).format('YYYY-MM-DD HH:mm:ss')}
                        </td>
                        <td>{item.user.name}</td>
                        <td>{item.user.email}</td>
                        <td>{item.customer.id ? 'Sim' : 'Não'}</td>
                        <td>
                          {String(item.payment.year + '-' + item.payment.month)}
                        </td>
                        <td>{item.paymentEmployee.name}</td>
                        <td>
                          {item.paymentEmployee.salary.toLocaleString('pt-br', {
                            style: 'currency',
                            currency: 'BRL',
                          })}
                        </td>
                      </tr>
                    ))
                  : ''}
              </tbody>
            </table>
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

export default withRouter(Alerts);
