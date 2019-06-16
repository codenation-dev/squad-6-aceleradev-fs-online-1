import 'filepond/dist/filepond.min.css';
import React, {Component} from 'react';
import {withRouter} from 'react-router';

import fileUploadCSV from './../services/UploadFileService';

class UploadCSV extends Component {
  constructor(props) {
    super(props);
    this.state = {
      file: null,
    };
  }

  handleFIle(e) {
    let file = e.target.files[0];
    this.setState({file: file});
  }

  async handleUpload(e) {
    e.preventDefault();
    fileUploadCSV.uploadFile(this.state.file);
  }

  render = () => (
    <div className="container">
      <div className="row">
        <h3>Upload de Clientes por arquivo CSV</h3>
        <br />
        <br />
      </div>
      <div className="row">
        <div className="input-group mb-3">
          <div className="input-group-prepend">
            <button className="btn btn-outline-secondary" type="button">
              Escolher Arquivo
            </button>
          </div>
          <input
            type="text"
            className="form-control"
            placeholder=""
            aria-label=""
            aria-describedby="basic-addon1"
          />
          <div className="input-group-append">
            <button className="btn btn-outline-secondary" type="button">
              Limpar
            </button>
            <button className="btn btn-outline-primary" type="button">
              Fazer Upload
            </button>
          </div>
        </div>

        <form>
          <div className="">
            <label>Selecione o arquivo</label>
            <input type="file" name="file" onChange={e => this.handleFIle(e)} />

            <button type="submit" onClick={e => this.handleUpload(e)}>
              Upload
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
export default withRouter(UploadCSV);
