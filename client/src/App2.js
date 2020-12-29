import React, { Component } from "react";
import Header from "./components/Header/Header";
import CssBaseline from "@material-ui/core/CssBaseline";
import RegistrationForm from "./components/RegistrationForm";
import { BrowserRouter as Router } from "react-router-dom";

class App2 extends Component {
  render() {
    return (
      <React.Fragment>
        <Router>
          <CssBaseline />
          <Header />
          <RegistrationForm />
        </Router>
      </React.Fragment>
    );
  }
}

export default App2;
