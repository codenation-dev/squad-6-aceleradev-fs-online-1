import loginService from './loginService';
const {CSV_FILE_UPLOAD} = require('./configApi');

const axios = require('axios');

async function uploadFile(file) {
  let files = file;

  let formdata = new FormData();
  formdata.append('file', files);

  const configRequest = {
    method: 'POST',
    headers: {
      Authorization: 'Bearer ' + loginService.userLogged().token,
    },
    url: CSV_FILE_UPLOAD,
    data: formdata,
  };

  try {
    const response = await axios(configRequest);
    if (response) {
      return response.data;
    }
  } catch (error) {}

  return false;
}

export default {uploadFile};
