import loginService from "./loginService";

const { GET_USERS_ENDPOINT } = require("./configApi");

const axios = require("axios");

async function getUsers() {
  const configRequest = {
    method: "GET",
    json: true,
    headers: {
      Authorization: "Bearer " + loginService.userLogged().token
    },
    url: GET_USERS_ENDPOINT
  };
  //efetua requisicao em si
  const response = await axios(configRequest);

  if (response) {
    return response.data;
  }
  return null;
}

export default { getUsers };
