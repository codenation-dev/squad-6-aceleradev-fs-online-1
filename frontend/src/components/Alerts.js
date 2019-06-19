import React, {Component} from 'react';
import {withRouter} from 'react-router';
import PropTypes from 'prop-types';

import alertService from '../services/alertService';

class Alerts extends Component {
  constructor(props) {
    super(props);
    this.state = {listAlerts: []};
  }

  async componentDidMount() {
    const retorno = await alertService.getAlerts();
    console.log(retorno,"Alert retorno")
    this.setState({
      listAlerts: retorno,
    });
  }

  btnNewClick(event) {
    this.props.history.push('/alert/');
    
  }

  btnEditClick(event, item) {
    this.props.history.push('/alert/' + item.id);
  }

  async btnDeleteClick(event, item) {
    await alertService.deleteAlert(item.id);
    const retorno = await alertService.getAlert();
    this.setState({
      listAlert: retorno,
    });
   }

 
  render = () => (
    <React.Fragment>
      <div className="container">
        <div className="row">
        {/*   <button
            onClick={e => {
              this.btnNewClick(e);
            }}
            type="button"
            className="btn btn-primary"
            id="btnNew"
            name="btnNew">            
            Novo Alerta
          </button> */}

        </div>
        <br />
        <div className="row">
          <table className="table table-hover">
            <thead>
              <tr>
                <th scope="col">#</th>
                <th scope="col">Arquivo</th>
                <th scope="col">Mês</th>
                <th scope="col">Ano</th>
                <th scope="col">Pagamento Nº</th>
              </tr>
            </thead>
            <tbody>
              {this.state.listAlerts
                ? this.state.listAlerts.map((item, index) => (

                    <tr key={item.id}>
                      <th scope="row">{item.id}</th>
                      <td>{item.filename}</td>
                      <td>{item.month}</td>
                      <td>{item.year}</td>
                      <td>{item.EmployeePayments}</td>
                      
                      <td>
                        <input
                          onClick={e => {
                            this.btnEditClick(e, item);
                          }}
                          className="btn btn-primary btn-sm"
                          type="button"
                          value="visualizar"
                          name="btnEdit"
                          id="btnEdit"
                        />
                    {/*     <input
                          onClick={e => {
                            this.btnDeleteClick(e, item);
                          }}
                          className="btn btn-danger btn-sm"
                          type="button"
                          value="Excluir"
                          name="btnDelete"
                          id="btnDelete"
                        /> */}
                      </td>
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
