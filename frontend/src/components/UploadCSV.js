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
<div class="input-group mb-3">

<div class="custom-file">
 <input type="file" class='custom-file-input'
 id='inputGroupFile03'   onChange={e => this.handleFIle(e)}
/>
  <label class="custom-file-label" 
  for="inputGroupFile03"  value={this.state.file} 

  >Escolha um arquivo</label>
  
</div>
<div className="input-group-append">
            <button className="btn btn-outline-secondary" type="button">
              Limpar
            </button>
            <button className="btn btn-outline-primary" type="button" onClick={e => this.handleUpload(e)}>
              Fazer Upload
            </button>
          </div>
</div>
</div>
  

   
    

     

        

        
      </div>
   
  );
}
export default withRouter(UploadCSV);
