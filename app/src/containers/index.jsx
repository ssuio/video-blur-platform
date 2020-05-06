import React, { useState, useEffect } from "react";
import { hot } from "react-hot-loader";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  Redirect,
} from "react-router-dom";
import VedioView from "./Vedio";
import { Login, Register, Logout } from "./User";
import Dashboard from "./Dashboard";
import Entry from "./Entry";
import Upload from './Upload';
import Pannel from './Pannel'
import apiHelper from "../services/api";

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
    <div>
      <div className="App">
        <Router>
          <Switch>
            <Route path="/entry" component={Entry} />
            <PrivateRoute path="/dashboard" component={Logout} />

            {/* <Route path="/dashboard/pannel" component={
              ()=> {return (
                <Dashboard>
                  <Pannel />
                </Dashboard>
              )}
            } />

            <Route path="/dashboard/upload" component={
              ()=> {return (
                <Dashboard>
                  <Upload />
                </Dashboard>
              )}
            } /> */}

            <Route path="*" exact={true} component={Page404} />
          </Switch>
        </Router>
      </div>
    </div>
  );
}

export default hot(module)(App);
