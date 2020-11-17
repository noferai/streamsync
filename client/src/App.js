import React, { Component } from "react";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput/ChatInput";
import { connect, sendMsg } from "./api";
import CssBaseline from "@material-ui/core/CssBaseline";
import Grid from "@material-ui/core/Grid";

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
        <CssBaseline />
        <Header />
        <Grid container>
          <Grid item xs={1} />
          <Grid item xs={11}>
            <ChatHistory chatHistory={this.state.chatHistory} />
            <ChatInput send={this.send} />
          </Grid>
        </Grid>
      </React.Fragment>
    );
  }
}

export default App;
