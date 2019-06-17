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
      lastModifiedDate:"c:/",
      type: "text/csv"
    
    }],
    arquivo :''
    };
   
  }

  handleFIle(e) {
     this.state.arquivo = e.target.files[0];
    this.setState({file: [{name: e.target.files[0].name,
    
    lastModified:e.target.files[0].lastModified,
    lastModifiedDate:e.target.files[0].lastModifiedDate,
    type: e.target.files[0].type
    }]});
      console.log(e.target.files[0],"daqui handlefile")
  }

  async handleUpload(e) {
    //e.preventDefault();
    fileUploadCSV.uploadFile(this.state.arquivo);
  }


  async handleDelete(e) {
     //e.preventDefault();
    this.setState({file: [{name: 'selecione o arquivo',
    lastModified:"1996",
    lastModifiedDate:"c:/",
    type: "text/csv"
      }]});
        console.log(this.state.file,"daqui do delete")
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
 id='inputGroupFile03' accept=".csv"   onChange={e => this.handleFIle(e)}
/>
  <label className="custom-file-label" 
  htmlfor="inputGroupFile03"  > {this.state.file[0].name}
{console.log(this.state.file)}

  </label>
  
</div>
<div className="input-group-append">
            <button className="btn btn-outline-secondary" type="button" 
            onClick={e => this.handleDelete(e)}
            
            disabled={this.state.file[0].name==="selecione o arquivo"}
            >
              Limpar
            </button>
            <button className="btn btn-outline-primary" type="button"
             onClick={e => this.handleUpload(e)}
             disabled={this.state.file[0].name==="selecione o arquivo"}
             >
              Fazer Upload
            </button>
          </div>
</div>
</div>
  

   
    

     

        

        
      </div>
   
  );
}
export default withRouter(UploadCSV);
