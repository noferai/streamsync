import React, { Component } from "react";
import "./ChatHistory.scss";
import Message from "../Message/Message";
import PropTypes from "prop-types";

class ChatHistory extends Component {
  render() {
    const messages = this.props.chatHistory.map((msg) => (
      <Message key={msg.timeStamp} message={msg.data} />
    ));

    return <div className="ChatHistory">{messages}</div>;
  }
}

ChatHistory.propTypes = {
  chatHistory: PropTypes.array,
};

export default ChatHistory;
