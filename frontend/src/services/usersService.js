import loginService from './loginService';

const {GET_USERS_ENDPOINT} = require('./configApi');

const axios = require('axios');

async function getUsers() {
  const configRequest = {
    method: 'GET',
    json: true,
    headers: {
      Authorization: 'Bearer ' + loginService.userLogged().token,
    },
    url: GET_USERS_ENDPOINT,
  };
  //efetua requisicao em si
  const response = await axios(configRequest);

  if (response) {
    return response.data;
  }
  return null;
}

async function getUserById(id) {
  const configRequest = {
    method: 'GET',
    json: true,
    headers: {
      Authorization: 'Bearer ' + loginService.userLogged().token,
    },
    url: GET_USERS_ENDPOINT + '/' + id,
  };
  //efetua requisicao em si
  const response = await axios(configRequest);

  if (response) {
    return response.data;
  }
  return null;
}

async function putUser(user) {
  const configRequest = {
    method: 'PUT',
    json: true,
    headers: {
      Authorization: 'Bearer ' + loginService.userLogged().token,
    },
    url: GET_USERS_ENDPOINT + '/' + user.id,
    //data: {...user, receiveAlert: String(user.receiveAlert)},
    data: user,
  };
  //efetua requisicao em si
  const response = await axios(configRequest);

  if (response) {
    return response.data;
  }
  return null;
}

async function postUser(user) {
  const configRequest = {
    method: 'POST',
    json: true,
    headers: {
      Accept: 'application/json',
      Authorization: 'Bearer ' + loginService.userLogged().token,
    },
    url: GET_USERS_ENDPOINT,
    data: user,
  };
  console.log(configRequest);
  //efetua requisicao em si
  const response = await axios(configRequest);

  if (response) {
    console.log(response);
    return response.data;
  }
  return null;
}

async function deleteUser(id) {
  const configRequest = {
    method: 'DELETE',
    json: true,
    headers: {
      Authorization: 'Bearer ' + loginService.userLogged().token,
    },
    url: GET_USERS_ENDPOINT + '/' + id,
  };
  //efetua requisicao em si
  const response = await axios(configRequest);

  if (response) {
    return response.data;
  }
  return null;
}

export default {getUsers, getUserById, putUser, postUser, deleteUser};
