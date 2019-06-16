import 'filepond/dist/filepond.min.css';
import React, {Component} from "react";
import {withRouter} from 'react-router';

import fileUploadCSV from "./../services/UploadFileService";



class UploadCSV extends Component {
    constructor(props) {
      super(props);
      this.state = {
        file:null
     
      };
    }
  
     handleFIle(e){
      let file = e.target.files[0]
      this.setState({file:file})
 
    }
   

    async handleUpload(e){
      e.preventDefault()
      fileUploadCSV.uploadFile(this.state.file)
     }
    render = () => (
        <div className="container">
        <form>
         <div className="">
          <label>Selecione o arquivo</label>
           <input type="file" name="file" onChange={(e) => 
            this.handleFIle(e)}/>

        <button type="submit" 
         onClick={(e)=>this.handleUpload(e)}
        >Upload </button>
         </div>
        </form>
       
                   
 
 
 </div>
  

  
    );
    }
  export default withRouter(UploadCSV);
  