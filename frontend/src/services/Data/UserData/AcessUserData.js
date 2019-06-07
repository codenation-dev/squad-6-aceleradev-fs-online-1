import {SERVER_URL,AUTH_TOKEN} from '../../Auth/JWTAuth' ;
import  React,{Component} from 'react';
import axios from 'axios';

class AcessUserData extends Component {

    state = {
        users: []
        
      }
      async getUser(data) {
       
        const GET_USER = `${SERVER_URL}/api/v1/user`;  
        //let response = await axios.post(GET_USER, data);
       const TOKEN =localStorage.getItem("acess_token");
       console.log(TOKEN)
       axios.defaults.headers.common = {'authorization': `bearer ${TOKEN}`}
        axios.get(GET_USER).then(response => response.data)
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