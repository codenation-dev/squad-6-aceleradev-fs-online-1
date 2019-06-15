import apiService from './apiService';

const {GET_USERS_ENDPOINT} = require('./configApi');
const endpoint = GET_USERS_ENDPOINT;

async function getUsers() {
  return await apiService.getApi(endpoint);
}

async function getUserById(id) {
  return apiService.getByIdApi(endpoint, id);
}

async function putUser(user) {
  return apiService.putApi(endpoint, user);
}

async function postUser(user) {
  return apiService.postApi(endpoint, user);
}

async function deleteUser(id) {
  return apiService.deleteApi(endpoint, id);
}

export default {getUsers, getUserById, putUser, postUser, deleteUser};
