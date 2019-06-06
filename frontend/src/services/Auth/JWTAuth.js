import axios from 'axios';
const SERVER_URL = "http://localhost:3000"

const login = async (data) => {

const LOGIN_ENDPOINT = `${SERVER_URL}/api/v1/signin`

try {
let response = await axios.post(LOGIN_ENDPOINT, data);

if (response.status === 200 && response.data.jwt && response.data.expireAt){

let jwt = response.data.jwt;
let expire_at = response.data.expireAt;

localStorage.setItem("acess_token", jwt);
localStorage.setItem("expires_at", expire_at);
console.log(jwt, "login");

}
}
catch(e){
console.log(e);

}
}

const register = async(data)=>{
  const SIGNUP_ENDEPOINT = `${SERVER_URL}/api/v1/signup`;
  
    try{
        let response = await axios({
          method :'post',
          responseType: 'json',
          url : SIGNUP_ENDEPOINT,
          data: data
        });

     }catch(e){
    console.log(e);

     }


 

}

const logout = () => {
    localStorage.removeItem("access_token");
    localStorage.removeItem("expire_at");
}


export {login, register , logout }