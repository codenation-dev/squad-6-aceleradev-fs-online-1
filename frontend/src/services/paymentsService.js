import http from './http';

const {GET_PAYMENTS_ENDPOINT} = require('./configApi');

async function getPayments(customerId) {
  const {data} = await http.get(
    GET_PAYMENTS_ENDPOINT + (customerId ? '?customerId=' + customerId : '')
  );

  return data;
}

async function getPaymentById(id) {
  const {data} = await http.get(`${GET_PAYMENTS_ENDPOINT}/${id}`);

  return data;
}

export default {
  getPayments,
  getPaymentById,
};
