<<<<<<< HEAD
import http from './http';

const {GET_USERS_ENDPOINT} = require('./configApi');

async function getUsers() {
  const {data} = await http.get(GET_USERS_ENDPOINT);

  return data;
}

async function getUserById(id) {
  const {data} = await http.get(`${GET_USERS_ENDPOINT}/${id}`);

  return data;
}

async function putUser(user) {
  const {data} = await http.put(`${GET_USERS_ENDPOINT}/${user.id}`, {user});

  return data;
}

async function postUser(user) {
  const {data} = await http.post(GET_USERS_ENDPOINT, {user});

  return data;
}

async function deleteUser(id) {
  const {data} = await http.delete(`${GET_USERS_ENDPOINT}/${id}`);

  return data;
=======
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
>>>>>>> 18213a1e9b2b325ba891429e53ed19e3c9ea5a20
}

export default {getUsers, getUserById, putUser, postUser, deleteUser};
