// import React from "react";
// import ReactDOM from "react-dom";
// // import Home from "./components/Home/Home";
// // import App from "./App";
// import App1 from "./App1";
// // // import App2 from "./App2";
// // // import App3 from "./App3";
// import * as serviceWorker from "./serviceWorker";
// //
// // ReactDOM.render(<App />, document.getElementById("root"));
// //
// ReactDOM.render(<App1 />, document.getElementById("root"));
// //
// // // ReactDOM.render(<App2 />, document.getElementById("root"));
// //
// // // ReactDOM.render(<App3 />, document.getElementById("root"));
// //
// // ReactDOM.render(<Home />, document.getElementById("root"));
// serviceWorker.unregister();

//////////////////////////////////////

import React from "react";
import ReactDOM from "react-dom";
// import { createBrowserHistory } from "history";
import { BrowserRouter as Router } from "react-router-dom";

import * as serviceWorker from "./serviceWorker";

// const history = createBrowserHistory();

import App4 from "./App4";

ReactDOM.render(
  <Router>
    <App4 />
  </Router>,
  document.getElementById("root")
);
serviceWorker.unregister();
