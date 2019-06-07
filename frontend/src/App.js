import React, {Component} from 'react';
import './App.css';
import {login, register,logout} from "./services/Auth/JWTAuth"

export default class App extends Component {
  async login(){
    let info = {
      email: "ruiblaese@gmail.com",
      password: "1234"
    };

    await login(info);

  }


  render(){
  return (
    <button className="btn estilo-1" onClick={this.login}> Login </button>
  );


}
}