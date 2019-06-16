import React from "react";
import PropTypes from "prop-types";
import { withRouter } from "react-router";

import loginService from "./../services/loginService";

const logoff = (e, p) => {
  loginService.logout();
  p.history.push("/login");
};

const Menu = props => (
  <nav className="navbar navbar-expand-lg navbar-light bg-light">
    <a className="navbar-brand" href="/">
      Banco Uati
    </a>
    <button
      className="navbar-toggler"
      type="button"
      data-toggle="collapse"
      data-target="#navbarSupportedContent"
      aria-controls="navbarSupportedContent"
      aria-expanded="false"
      aria-label="Toggle navigation"
    >
      <span className="navbar-toggler-icon" />
    </button>

    <div className="collapse navbar-collapse" id="navbarSupportedContent">
      <ul className="navbar-nav mr-auto">
        <li className="nav-item active">
          <a className="nav-link" href="/">
            Home <span className="sr-only">(current)</span>
          </a>
        </li>
        <li className="nav-item">
          <a className="nav-link" href="/users">
            Usuarios
          </a>
        </li>
        <li className="nav-item">
          <a className="nav-link disabled" href="/">
            Clientes
          </a>
        </li>
        <li className="nav-item">
          <a className="nav-link disabled" href="/">
            Pagamentos
          </a>
        </li>
        <li className="nav-item">
          <a className="nav-link disabled" href="/">
            Historico Alertas
          </a>
        </li>
       
      </ul>

      <button
        onClick={e => {
          logoff(e, props);
        }}
        className="btn btn-outline-secondary my-2 my-sm-0"
      >
        logoff
      </button>
    </div>
  </nav>
);

Menu.propTypes = {
  searchString: PropTypes.string,
  recipes: PropTypes.array
};

export default withRouter(Menu);
