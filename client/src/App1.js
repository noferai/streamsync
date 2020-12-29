import React, { Component } from "react";
import Header from "./components/Header/Header";
import CssBaseline from "@material-ui/core/CssBaseline";
import Chat from "./components/Chat";
import ReactPlayer from "react-player";
import { BrowserRouter as Router } from "react-router-dom";

class App1 extends Component {
  render() {
    return (
      <React.Fragment>
        <Router>
          <CssBaseline />
          <Header />
          <ReactPlayer
            controls="true"
            width="100%"
            height="75%"
            url="https://www.youtube.com/watch?v=ysz5S6PUM-U"
          />
          <Chat />
        </Router>
      </React.Fragment>
    );
  }
}

export default App1;
