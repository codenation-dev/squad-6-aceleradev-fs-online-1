import axios from 'axios';
const SERVER_URL = "http://localhost:3000";
const AUTH_TOKEN = '';
const login = async (data) => {

const LOGIN_ENDPOINT = `${SERVER_URL}/api/v1/signin`

try {
let response = await axios.post(LOGIN_ENDPOINT, data);
console.log(response.data.expire);
if (response.status === 200 && response.data.token && response.data.expire){

let jwt = response.data.token;
let expire_at = response.data.expire;

console.log("Teste Login !");
AUTH_TOKEN = localStorage.setItem("acess_token", jwt);
localStorage.setItem("expires_at", expire_at);

}
}
catch(e){

}
}

const registro = async(data)=>{
  const SIGNUP_ENDEPOINT = `${SERVER_URL}/api/v1/signup`;
  
    try{
      
        let response = await axios({
          method :'POST',
          headers: {            
            'Content-Type': 'application/json'
          },
          json: true,
          url : SIGNUP_ENDEPOINT,
          data: data
        });
       
     }catch(e){

        console.log(e);
     }

     console.log(data);
 

}

const logout = () => {
    localStorage.removeItem("access_token");
    localStorage.removeItem("expire_at");
    console.log("logout !")
}




  

export {login, registro , logout,SERVER_URL ,AUTH_TOKEN}