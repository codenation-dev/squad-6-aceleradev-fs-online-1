import 'filepond/dist/filepond.min.css';
import React, {Component} from "react";
import {withRouter} from 'react-router';
import axios from 'axios'
import {CSV_FILE_UPLOAD} from '../services/configApi'
import loginService from "../services/loginService";



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
/* 
console.log(e.target.files)
console.log(e.target.files[0]) */

 
    }
   

     handleUpload(e){
 

      let file = this.state.file
      let formdata = new FormData()
      formdata.append('image', file)
      formdata.append('name', 'clayton pereira')
     console.log(this.state.file,"handler upload")
 /*     axios({     
      method: 'POST',
      headers:{
        Authorization: "Bearer " + loginService.userLogged().token
     }, 
     url: CSV_FILE_UPLOAD,
    data:formdata
    }).then((res)=>{     
   
     },(err) => {

      console.log(err, "erro ao fazer upload")
     })
    */
     const configRequest = {
      method: "POST",     
      headers: {
        Authorization: "Bearer " + loginService.userLogged().token,
        'Content-Type': 'multipart/form-data'
      },
      url: CSV_FILE_UPLOAD,
      data:formdata
  };
    //efetua requisicao em si
    const response =  axios(configRequest);
    if (response) {
      return response.data;
    }
    return null;

     }
    render = () => (
        <div className="container">
        <form>
         <div className="">
          <label>Selecione o arquivo</label>
           <input type="file" name="file" onChange={(e) => 
            this.handleFIle(e)}/>

        <button type="button" 
         onClick={(e)=>this.handleUpload(e)}
        >Upload </button>
         </div>
        </form>
       
                   
 
 
 </div>
  

  
    );
    }
  export default withRouter(UploadCSV);
  