<<<<<<< HEAD
import http from './http';
import {localStorageWrapper} from '../helpers';

const NS_LOGGED_USER = 'logged_user';
const {SIGNUP_ENDPOINT} = require('./configApi');

async function login(email, password) {
  const {data} = await http.post(SIGNUP_ENDPOINT, {email, password});

  if (data) {
=======
import {localStorageWrapper} from '../helpers';

const axios = require('axios');

const NS_LOGGED_USER = 'logged_user';

const {SIGNUP_ENDPOINT} = require('./configApi');

async function login(username, password) {
  const configRequest = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    json: true,
    url: SIGNUP_ENDPOINT,
    data: JSON.stringify({
      email: username,
      password: password,
    }),
  };
  //efetua requisicao em si
  const response = await axios(configRequest);

  if (response) {
>>>>>>> 18213a1e9b2b325ba891429e53ed19e3c9ea5a20
    localStorageWrapper.set(NS_LOGGED_USER, {
      // DANGER: user credentials being stored in the local storage.
      email,
      password: password,
<<<<<<< HEAD
      token: data.token,
      expire: data.expire,
=======
      token: response.data.token,
      expire: response.data.expire,
>>>>>>> 18213a1e9b2b325ba891429e53ed19e3c9ea5a20
    });

    return true;
  }

  return false;
}

export const isLogged = () => !!localStorageWrapper.get(NS_LOGGED_USER);

export const userLogged = () => localStorageWrapper.get(NS_LOGGED_USER);

export const logout = () => localStorageWrapper.set(NS_LOGGED_USER, null);

export default {login, logout, isLogged, userLogged};
