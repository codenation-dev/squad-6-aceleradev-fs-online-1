import axios from 'axios';
const SERVER_URL = "http://localhost:3000"

const login = async (data) => {

const LOGIN_ENDPOINT = `${SERVER_URL}/api/v1/signin`

try {
let response = await axios.post(LOGIN_ENDPOINT, data);

if (response.status === 200 && response.data.jwt && response.data.expireAt){

let jwt = response.data.jwt;
let expire_at = response.data.expireAt


}




}catch(e){


}
}