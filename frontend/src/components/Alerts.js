import React, {Component} from 'react';
import {withRouter} from 'react-router';
import PropTypes from 'prop-types';

import alertService from '../services/alertService';

class Alerts extends Component {
  constructor(props) {
    super(props);
    this.state = {
      listAlerts: [],
    };
  }
  async componentDidMount() {
    const retorno = await alertService.getAlerts();
    console.log(retorno);
    this.setState({
      listAlerts: retorno,
    });
  }

  render = () => (
    <React.Fragment>
      <div className="container">
        <div className="row" />
        <br />
        <div className="row">
          <table className="table table-hover">
            <thead>
              <tr>
                <th scope="col">#</th>
                <th scope="col">Data\Hora</th>
                <th scope="col">Usuario Recebeu Alerta</th>
                <th scope="col">Usuario Email</th>
                <th scope="col">Cliente</th>
                <th scope="col">Pagamento</th>
                <th scope="col">Funcionario</th>
                <th scope="col">Salario</th>
              </tr>
            </thead>
            <tbody>
              {this.state.listAlerts
                ? this.state.listAlerts.map((item, index) => (
                    <tr key={item.id}>
                      <th scope="row">{item.id}</th>
                      <td>{item.date}</td>
                      <td>{item.user.name}</td>
                      <td>{item.user.email}</td>
                      <td>{item.customer.name}</td>
                      <td>{item.paymentEmployee.id}</td>
                      <td>{item.paymentEmployee.name}</td>
                      <td>{item.paymentEmployee.salary}</td>
                    </tr>
                  ))
                : ''}
            </tbody>
          </table>
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
