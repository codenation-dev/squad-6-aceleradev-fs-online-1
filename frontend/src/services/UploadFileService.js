import loginService from "./loginService";

const { CSV_FILE_UPLOAD } = require("./configApi");

const axios = require("axios");

async function uploadFile(file) {

    let files = file
    let formdata = new FormData()
  formdata.append('multipart/form-data', files)
  const configRequest = {
    method: "POST",
    json: true,
    headers: {
      Authorization: "Bearer " + loginService.userLogged().token
    },
    url: CSV_FILE_UPLOAD,
    data: formdata
  };
  //efetua requisicao em si
  const response = await axios(configRequest);

  if (response) {
    return response.data;
  }
  return null;
}

export default { uploadFile };
