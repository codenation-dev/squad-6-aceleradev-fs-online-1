import { FilePond } from 'react-filepond';
import 'filepond/dist/filepond.min.css';
import React, {Component} from "react";
import {withRouter} from 'react-router';

class UploadCSV extends Component {
    constructor(props) {
      super(props);
      this.state = {
     
      };
    }
 


  
   
    render = () => (
        <div className="container">
       
       <FilePond allowMultiple={false} 
       maxFiles={1} server='http://localhost:4000/api/v1/customer/upload'
       type/>
        </div>
  

  
    );
    }
  export default withRouter(UploadCSV);
  