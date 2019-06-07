import React, {Component} from 'react';
import './App.css';
import Login from './components/pages/Login/index';
import { login } from './services/Auth/JWTAuth';
import UserData from './services/Data/UserData/AcessUserData'

export default class App extends Component {
  constructor(props){
    super(props);
    this.state = {
         users: ''
    }
  }
    getUser(userData){

      this.setState({ users: userData })


    }
  

  
  render(){
  return (
   <div className="app-div">From app.js
   
   <Login/>
   <UserData ></UserData>
   <h1></h1>
   
   </div>
   
 
  );


}
}