import React, { Component } from "react";
import { withRouter } from "react-router";
import PropTypes from "prop-types";

import loginService from "../services/loginService";

class Login extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      password: ""
    };
  }
  handleChange(event) {
    this.setState({ [event.target.name]: event.target.value });
  }
  async login(event) {
    event.preventDefault();

    try {
      await loginService.login(this.state.username, this.state.password);
      this.props.history.push("/");
    } catch (error) {
      event.preventDefault();
      alert(error);
    }
  }

  render = () => (
    <form className="form-signin">
      <div className="text-center mb-4">
        <h1 className="h3 mb-3 font-weight-normal">Login / Register</h1>
      </div>

      <div className="form-label-group">
        <label htmlFor="inputEmail">Username</label>
        <input
          name="username"
          onChange={e => {
            this.handleChange(e);
          }}
          value={this.state.username}
          className="form-control"
          placeholder="Username"
          required
        />
      </div>

      <div className="form-label-group mt-2">
        <label htmlFor="inputPassword">Password</label>
        <input
          name="password"
          onChange={e => {
            this.handleChange(e);
          }}
          value={this.state.password}
          type="password"
          className="form-control"
          placeholder="Password"
          required
        />
      </div>

      <div className="mt-5">
        <button
          type="submit"
          onClick={e => {
            this.login(e);
          }}
          className="login btn btn-lg btn-primary btn-block"
        >
          Login
        </button>
      </div>
    </form>
  );
  static propTypes = {
    location: PropTypes.object.isRequired,
    history: PropTypes.object.isRequired
  };
}

export default withRouter(Login);
