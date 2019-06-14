import { localStorageWrapper } from "../helpers";

const axios = require("axios");

const NS_LOGGED_USER = "logged_user";

const { SIGNUP_ENDPOINT } = require("./configApi");

async function login(username, password) {
  const configRequest = {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    json: true,
    url: SIGNUP_ENDPOINT,
    data: JSON.stringify({
      email: username,
      password: password
    })
  };
  //efetua requisicao em si
  const response = await axios(configRequest);

  if (response) {
    localStorageWrapper.set(NS_LOGGED_USER, {
      email: username,
      password: password,
      token: response.data.token,
      expire: response.data.expire
    });
    return true;
  }
  return false;
}

export const isLogged = () => !!localStorageWrapper.get(NS_LOGGED_USER);

export const userLogged = () => localStorageWrapper.get(NS_LOGGED_USER);

export const logout = () => localStorageWrapper.set(NS_LOGGED_USER, null);

export default { login, logout, isLogged, userLogged };
