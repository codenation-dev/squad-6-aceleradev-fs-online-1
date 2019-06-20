import http from './http';

const {GET_ALERTS_ENDPOINT} = require('./configApi');

async function getAlerts(paramsQuery) {
  const {data} = await http.get(
    GET_ALERTS_ENDPOINT + (paramsQuery ? '?' + paramsQuery : '')
  );

  return data;
}

async function getAlertById(id) {
  const {data} = await http.get(`${GET_ALERTS_ENDPOINT}/${id}`);

  return data;
}

export default {
  getAlerts,
  getAlertById,
};
