import loginService from './loginService';

const axios = require('axios');

const getApi = async function(endpoint) {
  const configRequest = {
    method: 'GET',
    json: true,
    headers: {
      Authorization: 'Bearer ' + loginService.userLogged().token,
    },
    url: endpoint,
  };
  //efetua requisicao em si
  const response = await axios(configRequest);

  if (response) {
    return response.data;
  }
  return null;
};

async function getByIdApi(endpoint, id) {
  const configRequest = {
    method: 'GET',
    json: true,
    headers: {
      Authorization: 'Bearer ' + loginService.userLogged().token,
    },
    url: endpoint + '/' + id,
  };
  //efetua requisicao em si
  const response = await axios(configRequest);

  if (response) {
    return response.data;
  }
  return null;
}

async function putApi(endpoint, user) {
  const configRequest = {
    method: 'PUT',
    json: true,
    headers: {
      Authorization: 'Bearer ' + loginService.userLogged().token,
    },
    url: endpoint + '/' + user.id,
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

async function postApi(endpoint, user) {
  const configRequest = {
    method: 'POST',
    json: true,
    headers: {
      Accept: 'application/json',
      Authorization: 'Bearer ' + loginService.userLogged().token,
    },
    url: endpoint,
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

async function deleteApi(endpoint, id) {
  const configRequest = {
    method: 'DELETE',
    json: true,
    headers: {
      Authorization: 'Bearer ' + loginService.userLogged().token,
    },
    url: endpoint + '/' + id,
  };
  //efetua requisicao em si
  const response = await axios(configRequest);

  if (response) {
    return response.data;
  }
  return null;
}

export default {getApi, getByIdApi, putApi, postApi, deleteApi};
