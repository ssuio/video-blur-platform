import React, { useState, useEffect } from "react";
import { Redirect } from "react-router-dom";

import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';

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
                // return;
            })
            .catch(() => {
                setLoading(false);
            });
    };

    useEffect(checkLogin, []);

    function handleLogin() {
        apiHelper
            .login({account, passwd})
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
                        <TextField
                            label="User Name"
                            name="account"
                            onChange={onInputChange}
                            value={account}
                            type="text"
                            fullWidth={true}
                        />
                        <br /><br /><br />
                        <TextField
                            label="Password"
                            name="passwd"
                            onChange={onInputChange}
                            value={passwd}
                            type="password"
                            fullWidth={true}
                        />
                        <br /><br /><br /><br />
                        <Button variant="contained" fullWidth={true} onClick={handleLogin}>Login</Button>

                        <h6>I forget the password</h6>
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

export const Register = (props) => {
    const [account, setAccount] = useState("");
    const [passwd, setPasswd] = useState("");
    const [email, setEmail] = useState("");
    const [passwdAgain, setPasswdAgain] = useState("");

    const inputTable = {
        account: setAccount,
        passwd: setPasswd,
        email: setEmail,
        passwdAgain: setPasswdAgain
    };

    const onInputChange = (e) => {
        const name = e.target.name;
        const val = e.target.value;

        inputTable[name](val);
    };

    const handleRegister = () => {

        if(passwd !== passwdAgain){
            alert('Please ensure your password is right');
            return;
        }
        
        apiHelper.register({
            email:email,
            name:account,
            password:passwd
        })
        .then(()=>{
            props.goToLoginTab()
            alert("success")
        })
        .catch((err)=>{
            alert(err)
        })
    }

    return (
        <div>
            <TextField
                label="User Name"
                name="account"
                onChange={onInputChange}
                value={account}
                type="text"
                fullWidth={true}
            />
            <br /><br />

            <TextField
                label="Email"
                name="email"
                onChange={onInputChange}
                value={email}
                type="mail"
                fullWidth={true}
            />
            <br /><br />
            <TextField
                label="Password"
                name="passwd"
                onChange={onInputChange}
                value={passwd}
                type="password"
                fullWidth={true}
            />
            <br /><br />
            <TextField
                label="Password Again"
                name="passwdAgain"
                onChange={onInputChange}
                value={passwdAgain}
                type="password"
                fullWidth={true}
            />
            <br /><br /><br />
            <Button variant="contained" fullWidth={true} onClick={handleRegister}>Register</Button>
        </div>
    );
};
