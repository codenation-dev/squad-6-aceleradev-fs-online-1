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
  const {data} = await http.put(
    `${GET_CUSTOMERS_ENDPOINT}/${customer.id}`,
    customer
  );

  return data;
}

async function postCustomer(customer) {
  const {data} = await http.post(GET_CUSTOMERS_ENDPOINT, customer);

  return data;
}

async function deleteCustomer(id) {
  const {data} = await http.delete(`${GET_CUSTOMERS_ENDPOINT}/${id}`);

  return data;
}

export default {
  getCustomers,
  getCustomerById,
  putCustomer,
  postCustomer,
  deleteCustomer,
};
