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
  const response = await axios(configRequest);

  if (response) {
    alert("Arquivo CSV carregado com sucesso !")
    return response.data;
    
  }else{alert("Arquivo Não Carregado !")}

  return null;
}

export default {uploadFile};
