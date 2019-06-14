import React, { Component } from "react";
import { withRouter } from "react-router";
import PropTypes from "prop-types";

class Home extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render = () => (
    <React.Fragment>
      <div className="container">
        <div className="row">
          <h1>Banco Uati</h1>
          <p>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin
            imperdiet quam at leo tempor bibendum. Mauris sed ipsum nec nibh
            posuere sodales a eu urna. Nam bibendum lacus vitae ex scelerisque
            facilisis. Interdum et malesuada fames ac ante ipsum primis in
            faucibus. Vestibulum vitae orci eu eros pretium venenatis fringilla
            non magna. Phasellus tempor pulvinar enim sed fermentum. Proin
            rhoncus lorem a nibh egestas sollicitudin. Nulla posuere, neque ac
            viverra luctus, erat dolor blandit felis, a gravida erat mauris ut
            ligula. Praesent auctor felis enim, commodo vestibulum urna
            porttitor nec. Aliquam egestas vitae magna in ullamcorper. Nulla
            malesuada eros sed aliquet condimentum. Pellentesque turpis lacus,
            tincidunt sed lacus eget, imperdiet rutrum sem. Suspendisse eget
            nunc sed diam lacinia placerat a at est. Etiam lobortis blandit
            placerat. Duis ornare vestibulum erat ac finibus. Morbi tristique
            nisl mollis quam condimentum, in euismod velit varius. Aliquam eu
            mollis nibh. Aenean eget elit et quam egestas placerat ut facilisis
            tellus. Ut dictum nisl nec ligula mattis vehicula. Phasellus
            vehicula ligula at mauris luctus scelerisque. Proin posuere nulla
            commodo nunc cursus porta. Etiam pretium sed velit sit amet
            ultricies. Proin dictum aliquam ante a dapibus. Sed fringilla, elit
            vel mattis fermentum, ex augue luctus odio, at accumsan mi ante nec
            neque. Integer fringilla convallis mi, ut fermentum urna
            sollicitudin a. Maecenas sit amet tortor iaculis, ultricies elit ut,
            fermentum turpis. Integer sollicitudin, nulla quis ornare aliquam,
            ex risus semper risus, ut finibus ante libero a dolor.
          </p>
        </div>
      </div>
    </React.Fragment>
  );

  static propTypes = {
    location: PropTypes.object.isRequired,
    history: PropTypes.object.isRequired
  };
}

export default withRouter(Home);
