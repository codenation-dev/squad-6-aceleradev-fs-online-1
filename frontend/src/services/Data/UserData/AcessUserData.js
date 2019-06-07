import {SERVER_URL,AUTH_TOKEN} from '../../Auth/JWTAuth' ;
import  React,{Component} from 'react';
import axios from 'axios';

class AcessUserData extends Component {

    state = {
        users: []
        
      }
      
      
      getUser() {
        const axios = require('axios');
        axios.defaults.baseURL = SERVER_URL
        const completePath = `${SERVER_URL}/api/v1/user`;
        console.log("caminho_completo",completePath);

        axios.defaults.headers.common = {'Authorization': `bearer ${AUTH_TOKEN}`}
        const url = ''
        axios.get(url).then(response => response.data)
        .then((data) => {
          this.setState({ users: data })
          console.log("get users !")
         })
      }

      render(){
return(
    <div className="div-UserData">
<button onClick={this.getUser}> botao Get Users</button>
</div>
);

      }
}


export default AcessUserData;