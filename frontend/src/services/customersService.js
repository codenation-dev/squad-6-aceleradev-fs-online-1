import apiService from './apiService';

const {GET_CUSTOMERS_ENDPOINT} = require('./configApi');
const endpoint = GET_CUSTOMERS_ENDPOINT;

async function getCustomers() {
  return await apiService.getApi(endpoint);
}

async function getCustomerById(id) {
  return apiService.getByIdApi(endpoint, id);
}

async function putCustomer(Customer) {
  return apiService.putApi(endpoint, Customer);
}

async function postCustomer(Customer) {
  return apiService.postApi(endpoint, Customer);
}

async function deleteCustomer(id) {
  return apiService.deleteApi(endpoint, id);
}

export default {
  getCustomers,
  getCustomerById,
  putCustomer,
  postCustomer,
  deleteCustomer,
};
