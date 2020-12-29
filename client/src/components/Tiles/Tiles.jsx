import React, { Component } from "react";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import { Link } from "react-router-dom";
import { BrowserRouter as Router } from "react-router-dom";
import RoomImg from "../Home/room.png";

class Tiles extends Component {
  render() {
    const useStyles = makeStyles((theme) => ({
      root: {
        flexGrow: 1,
      },
      paper: {
        padding: theme.spacing(1),
        textAlign: "center",
        color: theme.palette.text.secondary,
      },
    }));

    function FormRow() {
      const classes = useStyles();
      return (
        <React.Fragment>
          <Router>
            <Grid item xs={4}>
              {/*<BrowserRouter>*/}
              <Link to={"/room"}>
                <Paper className={classes.paper}>
                  <img src={RoomImg} width="100" height="111" />
                  <div>Room</div>
                </Paper>
              </Link>
              {/*</BrowserRouter>*/}
            </Grid>
            <Grid item xs={4}>
              {/*<BrowserRouter>*/}
              <Link to={"/room"}>
                <Paper className={classes.paper}>
                  <img src={RoomImg} width="100" height="111" />
                  <div>Room</div>
                </Paper>
              </Link>
              {/*</BrowserRouter>*/}
            </Grid>
            <Grid item xs={4}>
              {/*<BrowserRouter>*/}
              <Link to={"/room"}>
                <Paper className={classes.paper}>
                  <img src={RoomImg} width="100" height="111" />
                  <div>Room</div>
                </Paper>
              </Link>
              {/*</BrowserRouter>*/}
            </Grid>
          </Router>
        </React.Fragment>
      );
    }

    return (
      <div>
        <Grid container spacing={1}>
          <Grid container item xs={12} spacing={3}>
            <FormRow />
          </Grid>
          <Grid container item xs={12} spacing={3}>
            <FormRow />
          </Grid>
          <Grid container item xs={12} spacing={3}>
            <FormRow />
          </Grid>
        </Grid>
      </div>
    );
  }
}
export default Tiles;
