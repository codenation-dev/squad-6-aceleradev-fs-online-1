import React, { Component } from 'react';
import { withRouter } from 'react-router';
import PropTypes from 'prop-types';

import alertService from '../services/alertService';

class Alerts extends Component {
  constructor(props) {
    super(props);
    this.state = {
      listAlerts: [],
      searchString: ''
    };

  }
  async componentDidMount() {
    const retorno = await alertService.getAlerts();
    console.log(retorno, "Alert retorno")
    this.setState({
      listAlerts: [retorno[0]],
      searchString: '',
    });
  }

  filtraValor(e) {
    // Variable to hold the original version of the list
    let currentList = [];
    // Variable to hold the filtered list before putting into state
    let newList = [];

    
    if (e.target.value !== "") {
      if(this.state.listAlerts){
   
      currentList = this.state.listAlerts
      console.log(currentList,"alertas")
      }
   
      newList = currentList.filter(item => {
       
        const lc = item.toLowerCase();
       
        const filter = e.target.value.toLowerCase();
       
        return lc.includes(filter);
      });
    } else {
 
      newList = this.props.items;
    }
    
    this.setState({
      listAlerts: newList
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
          <input type="text"
            name="pesquisa"
            onChange={(e) => this.filtraValor(e)} />


          <button >  pesquisar </button>

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
