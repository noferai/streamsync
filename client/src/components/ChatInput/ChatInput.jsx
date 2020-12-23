import React, { Component } from "react";
import "./ChatInput.scss";
import PropTypes from "prop-types";

class ChatInput extends Component {
  render() {
    return (
      <div className="ChatInput">
        <input
          onKeyDown={this.props.send}
          placeholder="Ваше сообщение здесь..."
        />
      </div>
    );
  }
}

ChatInput.propTypes = {
  send: PropTypes.func,
};

export default ChatInput;
