import React, {Component} from 'react';
import {withRouter} from 'react-router';
import PropTypes from 'prop-types';

import loginService from '../services/loginService';

class Login extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: '',
      password: '',
    };
  }
  handleChange(event) {
    this.setState({[event.target.name]: event.target.value});
  }
  async login(event) {
    event.preventDefault();

    try {
      await loginService.login(this.state.username, this.state.password);
      this.props.history.push('/');
    } catch (error) {
      event.preventDefault();
      if (String(error).indexOf('401') >= 0) {
        alert('Usuario ou/e senha incorreto(s).');
      } else {
        alert('Erro ao tentar fazer o login.\n' + error);
      }
    }
  }

  render = () => (
    <div className="container">
      <br />
      <div className="text-center mb-4">
        <h1 className="h3 mb-3 font-weight-normal">Login</h1>
      </div>

      <form className="form-signin">
        <div className="form-label-group">
          <label htmlFor="inputEmail">E-mail</label>
          <input
            name="username"
            onChange={e => {
              this.handleChange(e);
            }}
            value={this.state.username}
            className="form-control"
            placeholder="e-mail"
            required
          />

          <div className="form-label-group mt-2">
            <label htmlFor="inputPassword">Senha</label>
            <input
              name="password"
              onChange={e => {
                this.handleChange(e);
              }}
              value={this.state.password}
              type="password"
              className="form-control"
              placeholder="senha"
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
        </div>
      </form>
    </div>
  );
  static propTypes = {
    location: PropTypes.object.isRequired,
    history: PropTypes.object.isRequired,
  };
}

export default withRouter(Login);
