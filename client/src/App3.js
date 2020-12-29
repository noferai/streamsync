import React, { Component } from "react";
import Header from "./components/Header/Header";
import Chat from "./components/Chat/Chat";
// import ChatInput from "./components/ChatInput/ChatInput";
import ReactPlayer from "react-player";
import { BrowserRouter as Router } from "react-router-dom";
import { connect, sendMsg } from "./api";
import CssBaseline from "@material-ui/core/CssBaseline";
import Grid from "@material-ui/core/Grid";

import Tiles from "./components/Tiles";
import RegistrationForm from "./components/RegistrationForm";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
    };
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message");
      this.setState((prevState) => ({
        chatHistory: [...prevState.chatHistory, msg],
      }));
      console.log(this.state);
    });
  }

  send(event) {
    if (event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }

  render() {
    return (
      <React.Fragment>
        <Router>
          <CssBaseline />
          <Header />
          <Grid container>
            <Grid item xs={1} />
            <Grid item xs={11}></Grid>
          </Grid>
          <Tiles></Tiles>
          <ReactPlayer
            controls="true"
            width="100%"
            height="75%"
            url="https://www.youtube.com/watch?v=ysz5S6PUM-U"
          />
          <Chat />
          <RegistrationForm />
        </Router>
      </React.Fragment>
    );
  }
}

export default App;
