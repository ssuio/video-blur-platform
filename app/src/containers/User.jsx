import React, { useState, useEffect } from "react";
import { Redirect } from "react-router-dom";
import apiHelper from "../services/api";

export const LoginView = () => {
  const [loading, setLoading] = useState(true);
  const [user, setUser] = useState({});
  const [account, setAccount] = useState("");
  const [passwd, setPasswd] = useState("");

  return (
    <div>
      {!loading && user.email ? (
        <Login />
      ) : (
        <Redirect
          to={{ pathname: "/dashboard", state: { from: props.location } }}
        />
      )}
    </div>
  );
};

export const Login = (props) => {
  const [loading, setLoading] = useState(true);
  const [user, setUser] = useState({});
  const [account, setAccount] = useState("");
  const [passwd, setPasswd] = useState("");

  const inputTable = {
    account: setAccount,
    passwd: setPasswd,
  };

  const checkLogin = () => {
    apiHelper
      .profile()
      .then((user) => {
        setLoading(false);
        setUser(user);
        return;
      })
      .catch(() => {
        setLoading(false);
      });
  };

  useEffect(checkLogin, []);

  function login() {
    apiHelper
      .login(account, passwd)
      .then(() => apiHelper.profile())
      .then((user) => {
        console.log(user);
        setUser(user);
      })
      .catch(console.error);
  }

  const onInputChange = (e) => {
    const name = e.target.name;
    const val = e.target.value;

    inputTable[name](val);
  };

  return (
    <div>
      {!loading && user.email ? (
        <Redirect to={{ pathname: "/dashboard", state: { user } }} />
      ) : (
        <div>
          <h1>Login</h1>
          <br />
          <label>Account</label>
          <input
            name="account"
            onChange={onInputChange}
            value={account}
            type="text"
          />
          <br />
          <label>Password</label>
          <input
            name="passwd"
            onChange={onInputChange}
            value={passwd}
            type="password"
          />
          <br />
          <button onClick={login}>Login</button>
        </div>
      )}
    </div>
  );
};

export const Logout = () => {
  function logout() {
    apiHelper
      .logout()
      .then(() => {
      })
      .catch(console.error);
  }

  return <button onClick={logout}>Logout</button>;
};

export const Register = () => {
  return (
    <div>
      <h1> Regeister </h1>
    </div>
  );
};
