<<<<<<< HEAD
import http from './http';

const {GET_CUSTOMERS_ENDPOINT} = require('./configApi');

async function getCustomers() {
  const {data} = await http.get(GET_CUSTOMERS_ENDPOINT);

  return data;
}

async function getCustomerById(id) {
  const {data} = await http.get(`${GET_CUSTOMERS_ENDPOINT}/${id}`);

  return data;
}

async function putCustomer(customer) {
  const {data} = await http.put(`${GET_CUSTOMERS_ENDPOINT}/${customer.id}`, {
    customer,
  });

  return data;
}

async function postCustomer(customer) {
  const {data} = await http.post(GET_CUSTOMERS_ENDPOINT, {customer});

  return data;
}

async function deleteCustomer(id) {
  const {data} = await http.delete(`${GET_CUSTOMERS_ENDPOINT}/${id}`);

  return data;
=======
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
>>>>>>> 18213a1e9b2b325ba891429e53ed19e3c9ea5a20
}

export default {
  getCustomers,
  getCustomerById,
  putCustomer,
  postCustomer,
  deleteCustomer,
};
