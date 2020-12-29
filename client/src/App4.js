import React, { Component } from "react";

import { Route } from "react-router";
import { Router } from "react-router";
import { Switch } from "react-router";
import { Redirect } from "react-router-dom";
import { withRouter } from "react-router-dom";
// import Home from "./components/Home/Home";

import App from "./App";
import App1 from "./App1";
import App2 from "./App2";

class App4 extends Component {
  render() {
    return (
      <Router>
        <div className="App">
          <Switch>
            <Route path="/main" component={App} />
            <Route path="/room" component={App1} />
            <Route path="/registration" component={App2} />
            <Redirect from="/" to="/main" />
          </Switch>
        </div>
      </Router>
    );
  }
}

export default withRouter(App4);
