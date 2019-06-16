import React, {Component} from 'react';
import {withRouter} from 'react-router';
import PropTypes from 'prop-types';

import usersService from './../services/usersService';
import loginService from '../services/loginService';

class User extends Component {
  constructor(props) {
    super(props);
    this.state = {
      userInvalid: false,
      userId: 0,
      userForm: {
        id: 0,
        email: '',
        password: '',
        name: '',
        receiveAlert: false,
      },
    };
  }

  async componentDidMount() {
    if (this.props.match.params.id) {
      let user;
      try {
        user = await usersService.getUserById(this.props.match.params.id);
      } catch (error) {
        console.log(error);
      }
      if (user) {
        this.setState({
          userId: this.props.match.params.id,
          userForm: user,
        });
      } else {
        this.setState({userInvalid: true});
      }
    }
  }

  btnSalvarClick = async event => {
    event.preventDefault();

    let msg = '';
    if (!this.state.userForm.email) {
      msg += 'Preencha o campo Email\n';
    }
    if (!this.state.userForm.password) {
      msg += 'Preencha o campo Senha\n';
    }
    if (!this.state.userForm.name) {
      msg += 'Preencha o campo Nome\n';
    }

    if (!msg) {
      if (this.state.userForm.id) {
        await usersService.putUser(this.state.userForm);
        this.props.history.push('/users');
      } else {
        await usersService.postUser(this.state.userForm);
        this.props.history.push('/users');
      }
    } else {
      alert(msg);
    }
  };

  btnCancelarClick = event => {
    event.preventDefault();
    this.props.history.push('/users');
  };

  handleOnChange(event) {
    if (event.target.name !== 'receiveAlert') {
      this.setState({
        ...this.state,
        userForm: {
          ...this.state.userForm,
          [event.target.name]: event.target.value ? event.target.value : '',
        },
      });
    } else {
      this.setState({
        ...this.state,
        userForm: {
          ...this.state.userForm,
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
                value={this.state.userId}
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
                value={this.state.userForm.email}
                onChange={e => {
                  this.handleOnChange(e);
                }}
                required
              />
            </div>
            <div className="form-group col-md-3">
              <label htmlFor="inputPassword">Senha</label>
              <input
                disabled={this.state.userId !== 0}
                type="password"
                className="form-control"
                placeholder="Password"
                onChange={e => {
                  this.handleOnChange(e);
                }}
                id="inputPassword"
                name="password"
                value={this.state.userForm.password}
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
              value={this.state.userForm.name}
            />
          </div>

          <div className="form-group">
            <div className="form-check">
              <input
                className="form-check-input"
                type="checkbox"
                onChange={e => {
                  this.handleOnChange(e);
                }}
                id="gridCheck"
                name="receiveAlert"
                checked={this.state.userForm.receiveAlert}
              />
              <label className="form-check-label" htmlFor="gridCheck">
                Receber Alertas?
              </label>
            </div>
          </div>
          <div className="form-group">
            <button
              disabled={this.state.userInvalid}
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

export default withRouter(User);
