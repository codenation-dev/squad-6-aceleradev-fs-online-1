import {SERVER_URL,AUTH_TOKEN} from '../../Auth/JWTAuth' ;
import  React,{Component} from 'react';
import axios from 'axios';

class AcessUserData extends Component {

    state = {
        users: []
        
      }
      async getUser(data) {
       data.persist()
        try{
          const GET_USER = `${SERVER_URL}/api/v1/user`;  
          const TOKEN =localStorage.getItem("acess_token");
       
          let response = await axios({
            method :'GET',
            headers: {            
              'Content-Type': 'application/json',
              'Authorization': `bearer ${TOKEN}`
            },
            json: true,
            url : GET_USER,
            //data: data
          });
         
       }catch(e){
  
          console.log(e);
       }
  
       console.log(data);
   
  
  }
        /* const TOKEN =localStorage.getItem("acess_token");
      
      
        axios.defaults.headers.common = {'authorization': `bearer ${TOKEN}`}
        const GET_USER = `${SERVER_URL}/api/v1/user`;  
        //let response = await axios.post(GET_USER, data);
       
        axios.get(GET_USER).then(response => response.data)
        .then((data) => {
          this.setState({ users: data })
          console.log( "dados do resquest")

         })

 
}*/
      

      
      render(){
return(
    <div className="div-UserData">
<button onClick={this.getUser}> botao Get Users</button>
</div>
);

      }
}


export default AcessUserData;