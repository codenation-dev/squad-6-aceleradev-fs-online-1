import React,{Component} from 'react';
import {login, registro,logout} from '../../../../src/services/Auth/JWTAuth'

class Login extends Component {

    async login(){
        let info = {
          email: "ruiblaese@gmail.com",
          password: "1234"
        };
    
         await login(info);
        
      }


render(){
return( 
<div className="buttons">
<h1 className="h1"> Tela de Login</h1>
<button className="btn estilo-1" onClick={this.login}> Login </button>
 
{/* <button className="btn btn-primary" onClick = {logout }>Log out</button> */}
<button className="btn btn-primary" onClick = {registro}>Registre-se</button>
</div>
);

}
}

export default Login;