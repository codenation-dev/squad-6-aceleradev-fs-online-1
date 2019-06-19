import http from './http';

const {GET_ALERTS_ENDPOINT} = require('./configApi');

async function getAlerts() {
  const {data} = await http.get(GET_ALERTS_ENDPOINT);

  return data;
}

async function getAlertById(id) {
  const {data} = await http.get(`${GET_ALERTS_ENDPOINT}/${id}`);

  return data;
}

async function putAlert(alert) {
  const {data} = await http.put(
    `${GET_ALERTS_ENDPOINT}/${alert.id}`,
    alert
  );

  return data;
}

async function postAlert(alert) {
  const {data} = await http.post(GET_ALERTS_ENDPOINT, alert);

  return data;
}

async function deleteAlert(id) {
  const {data} = await http.delete(`${GET_ALERTS_ENDPOINT}/${id}`);

  return data;
}

export default {
  getAlerts,
  getAlertById,
  putAlert,
  postAlert,
  deleteAlert,
};
