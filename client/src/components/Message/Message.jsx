import React, {Component} from 'react';
import './Message.scss';
import Typography from "@material-ui/core/Typography";

class Message extends Component {
    constructor(props) {
        super(props);
        let temp = JSON.parse(this.props.message);
        this.state = {
            message: temp
        }
    }

    render() {
        return (
            <Typography className="Message" variant="body2">
                    {this.state.message.body}
            </Typography>
        );
    };

}

export default Message;