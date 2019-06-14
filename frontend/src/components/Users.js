import React, { Component } from "react";
import { withRouter } from "react-router";
import PropTypes from "prop-types";

import usersService from "./../services/usersService";

class Users extends Component {
  constructor(props) {
    super(props);
    this.state = { listUsers: [] };
  }

  async componentDidMount() {
    const retorno = await usersService.getUsers();
    console.log(retorno);
    this.setState({
      listUsers: retorno
    });
  }

  render = () => (
    <React.Fragment>
      <div className="container">
        <table className="table table-hover">
          <thead>
            <tr>
              <th scope="col">#</th>
              <th scope="col">Email</th>
              <th scope="col">Nome</th>
              <th scope="col">Recebe Aleras?</th>
            </tr>
          </thead>
          <tbody>
            {this.state.listUsers
              ? this.state.listUsers.map((item, index) => (
                  <tr key={item.id}>
                    <th scope="row">{item.id}</th>
                    <td>{item.email}</td>
                    <td>{item.name}</td>
                    <td>{item.receiveAlert ? "Sim" : "Nao"}</td>
                  </tr>
                ))
              : ""}
          </tbody>
        </table>
      </div>
    </React.Fragment>
  );

  static propTypes = {
    location: PropTypes.object.isRequired,
    history: PropTypes.object.isRequired
  };
}

export default withRouter(Users);
