import React, { Component } from "react";
import "./Header.scss";
import AppBar from "@material-ui/core/AppBar";
import Button from "@material-ui/core/Button";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import { Link } from "react-router-dom";
import { BrowserRouter } from "react-router-dom";

class Header extends Component {
  render() {
    return (
      <AppBar position="sticky" elevation={0}>
        <Toolbar className="header">
          <Typography variant="h6" color="inherit" noWrap>
            <BrowserRouter>
              <Link to={"/exit"}>
                <Button variant="contained">Exit</Button>
              </Link>
              <Link to={"/main"}>
                <Button variant="contained" color="primary">
                  View All Rooms
                </Button>
              </Link>
              <Link to={"/registration"}>
                <Button variant="contained" color="primary">
                  Authentication
                </Button>
              </Link>
            </BrowserRouter>
            <div>Avtomatgorizonta</div>
          </Typography>
        </Toolbar>
      </AppBar>
    );
  }
}

export default Header;
