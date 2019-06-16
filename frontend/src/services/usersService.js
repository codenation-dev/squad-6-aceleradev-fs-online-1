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
  const {data} = await http.put(`${GET_USERS_ENDPOINT}/${user.id}`, user);

  return data;
}

async function postUser(user) {
  const {data} = await http.post(GET_USERS_ENDPOINT, user);

  return data;
}

async function deleteUser(id) {
  const {data} = await http.delete(`${GET_USERS_ENDPOINT}/${id}`);

  return data;
}

export default {getUsers, getUserById, putUser, postUser, deleteUser};
