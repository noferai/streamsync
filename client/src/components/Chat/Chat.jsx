import React, { Component } from "react";
import "./Chat.scss";
import PropTypes from "prop-types";

const DATA = [
  {
    senderId: "Ernest",
    text: "LOL?",
  },
  {
    senderId: "Max",
    text: "KEK!",
  },
];

// eslint-disable-next-line no-unused-vars
const testToken =
  "https://us1.pusherplatform.io/services/chatkit_token_provider/v1/dfaf1e22-2d33-45c9-b4f8-31f634621d24/token";
// eslint-disable-next-line no-unused-vars
const instanceLocator = "v1:us1:dfaf1e22-2d33-45c9-b4f8-31f634621d24";
const roomId = 1;
// const username = "perborgen";

class Chat extends Component {
  constructor(props) {
    super(props);
    this.state = {
      messages: DATA,
    };
    this.sendMessage = this.sendMessage.bind(this);
  }

  // componentDidMount() {
  //   const chatManager = new Chatkit.ChatManager({
  //     instanceLocator: instanceLocator,
  //     userId: "janedoe",
  //     // eslint-disable-next-line no-undef
  //     tokenProvider: new Chatkit.TokenProvider({
  //       url: testToken,
  //     }),
  //   });
  //
  //   chatManager.connect().then((currentUser) => {
  //     this.currentUser = currentUser;
  //     this.currentUser.subscribeToRoom({
  //       roomId: roomId,
  //       hooks: {
  //         onNewMessage: (message) => {
  //           this.setState({
  //             messages: [...this.state.messages, message],
  //           });
  //         },
  //       },
  //     });
  //   });
  // }

  sendMessage(text) {
    this.currentUser.sendMessage({
      text,
      roomId: roomId,
    });
  }

  render() {
    return (
      <div className="app">
        <Title />
        <MessageList
          roomId={this.state.roomId}
          messages={this.state.messages}
        />
        <SendNickname sendNickname={this.sendMessage} />
        <SendMessageForm sendMessage={this.sendMessage} />
      </div>
    );
  }
}

class MessageList extends Component {
  render() {
    return (
      <ul className="message-list">
        {/* eslint-disable-next-line no-unused-vars */}
        {this.props.messages.map((message, index) => {
          return (
            <li key={message.id} className="message">
              <div>{message.senderId}</div>
              <div>{message.text}</div>
            </li>
          );
        })}
      </ul>
    );
  }
}

class SendMessageForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      message: "",
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(e) {
    this.setState({
      message: e.target.value,
    });
  }

  handleSubmit(e) {
    e.preventDefault();
    this.props.sendMessage(this.state.message);
    this.setState({
      message: "",
    });
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit} className="send-message-form">
        <input
          onChange={this.handleChange}
          value={this.state.message}
          placeholder="Please type your message"
          type="text"
        />
      </form>
    );
  }
}

class SendNickname extends Component {
  constructor(props) {
    super(props);
    this.state = {
      message: "",
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(e) {
    this.setState({
      message: e.target.value,
    });
  }

  handleSubmit(e) {
    e.preventDefault();
    this.props.sendMessage(this.state.message);
    this.setState({
      message: "",
    });
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit} className="send-message-form">
        <input
          onChange={this.handleChange}
          value={this.state.message}
          placeholder="Please type your nickname"
          type="text"
        />
      </form>
    );
  }
}

class Title extends Component {
  render() {
    return <p className="title">Chat</p>;
  }
}

Chat.propTypes = {
  sendMessage: PropTypes.func,
};

SendMessageForm.propTypes = {
  sendMessage: PropTypes.func,
};

SendNickname.propTypes = {
  sendMessage: PropTypes.func,
};

MessageList.propTypes = {
  messages: PropTypes.func,
};

export default Chat;
