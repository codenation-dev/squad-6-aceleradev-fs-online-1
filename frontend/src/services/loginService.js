import http from './http';
import {localStorageWrapper} from '../helpers';

const NS_LOGGED_USER = 'logged_user';
const {SIGNUP_ENDPOINT} = require('./configApi');

async function login(email, password) {
  const {data} = await http.post(SIGNUP_ENDPOINT, {email, password});

  if (data) {
    localStorageWrapper.set(NS_LOGGED_USER, {
      // DANGER: user credentials being stored in the local storage.
      email,
      password: password,
      token: data.token,
      expire: data.expire,
    });

    return true;
  }

  return false;
}

export const isLogged = () => !!localStorageWrapper.get(NS_LOGGED_USER);

export const userLogged = () => localStorageWrapper.get(NS_LOGGED_USER);

export const logout = () => localStorageWrapper.set(NS_LOGGED_USER, null);

export default {login, logout, isLogged, userLogged};
