import axios from 'axios';
import loginService from './loginService';

const http = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL,
  timeout: 3000,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
  },
});

http.interceptors.request.use(
  async config => {
    const token = loginService.userLogged() && loginService.userLogged().token;

    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }

    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

http.interceptors.response.use(
  response => response,
  async error => {
    if (error.response === undefined) {
      console.warn('The API server is down');
    }

    if (error.response && error.response.status === 401) {
      loginService.logout();
    } else if (error.response && error.response.status === 403) {
      loginService.logout();
    }

    return Promise.reject(error);
  }
);

export default http;
