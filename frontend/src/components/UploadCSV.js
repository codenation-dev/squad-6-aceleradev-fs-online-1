import 'filepond/dist/filepond.min.css';
import React, {Component} from 'react';
import {withRouter} from 'react-router';

import fileUploadCSV from './../services/UploadFileService';

class UploadCSV extends Component {
  constructor(props) {
    super(props);
    this.state = {
      file: [{ name: 'selecione o arquivo',
      lastModified:"1996",
      lastModifiedDate:"c:/"
    
    }]
    };
  }

  handleFIle(e) {
    let file = e.target.files[0];
    this.setState({file: file});
    console.log(this.state.file)
  }

  async handleUpload(e) {
    //e.preventDefault();
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

<div className="custom-file">
 <input type="file" className='custom-file-input'
 id='inputGroupFile03'    onChange={e => this.handleFIle(e)}
/>
  <label className="custom-file-label" 
  htmlfor="inputGroupFile03"   > {this.state.file[0].name}
{console.log(this.state.file[0].name)}

  </label>
  
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
