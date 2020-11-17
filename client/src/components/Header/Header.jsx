import React, { Component } from "react";
import "./Header.scss";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";

class Header extends Component {
  render() {
    return (
      <AppBar position="sticky" elevation={0}>
        <Toolbar className="header">
          <Typography variant="h6" color="inherit" noWrap>
            Avtomatgorizonta
          </Typography>
        </Toolbar>
      </AppBar>
    );
  }
}

export default Header;
