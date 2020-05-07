import React, { useState, useEffect } from "react";
import { hot } from "react-hot-loader";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
} from "react-router-dom";
import Entry from "./Entry";
import apiHelper from "../services/api";
import {Dashboard} from "./Dashboard";

const Page404 = () => <h1> Page 404 </h1>;
const Loading = () => <h1>Loading ... </h1>;

const PrivateRoute = ({ component: Component, ...rest }) => {
  const [loading, setLoading] = React.useState(true);
  const [auth, setAuth] = React.useState(false);

  const checkAuth = () => {
    apiHelper.profile()
        .then(() => {
          console.log("Succ auth");
          setAuth(true);
          setLoading(false);
        })
        .catch(()=>{
          console.log("Failed auth");
          setLoading(false);
        });
  };

  useEffect(checkAuth, []);

  return (
    <Route
      {...rest}
      render={(props) =>
        loading ? (
          <Loading />
        ) : auth ? (
          <Component {...props} />
        ) : (
          <Redirect
            to={{ pathname: "/login", state: { from: props.location } }}
          />
        )
      }
    />
  );
};

function App() {
  return (
    <div className="App">
        <Router>
          <Switch>
            <Route path="/entry" component={Entry} />
            <PrivateRoute path="/dashboard" component={Dashboard} />
            <Route path="*" exact={true} component={Page404} />
          </Switch>
        </Router>
    </div>
  );
}

export default hot(module)(App);
