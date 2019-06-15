import 'filepond/dist/filepond.min.css';
import React, {Component} from "react";
import {withRouter} from 'react-router';

class UploadCSV extends Component {
    constructor(props) {
      super(props);
      this.state = {
        file:null
     
      };
    }
  
    handleFIle(e){
      let file
      this.setState({file:file})
/* 
console.log(e.target.files)
console.log(e.target.files[0]) */


    }
   

     handleUpload(e){

     console.log()


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
  