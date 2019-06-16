import loginService from './loginService';
import {localStorageWrapper} from '../helpers';
const {CSV_FILE_UPLOAD} = require('./configApi');

const axios = require('axios');

async function uploadFile(file) {
  let files = file;
  console.log(files, 'dentro da função de upload!');

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

  //efetua requisicao em si
  console.log(formdata, 'arquivo formdata');
  console.log(
    'Bearer ' + loginService.userLogged().token,
    'token from login !'
  );
  console.log(
    'Bearer ' + localStorageWrapper.get('logged_user').token,
    'token from storage !'
  );
  console.log(configRequest, 'here request');

  const response = await axios(configRequest);

  if (response) {
    return response.data;
  }
  return null;
}

export default {uploadFile};
