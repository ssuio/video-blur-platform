import React, { useState } from "react";
import { Redirect } from "react-router-dom";
import icLogoFull from "assets/images/ic_logo_full.png";
import apiHelper from "../services/api";

const FinishPopup = (props) => {
    return (
        <div>
            <div className="popup">
                <div className="popupMessageContainer">
                    <div className="popupContentContainer">
                        <div className="popupContentTitle">Success !</div>
                        Please login by your email and password.
                    </div>
                    <div className="popupBtnContainer">
                        <button className="btnStyle sizeL green" onClick={props.clostPopup}>Close</button>
                    </div>
                </div>
            </div>
        </div>
    );
};

const Login = () => {
    const [account, setAccount] = useState("");
    const [passwd, setPasswd] = useState("");
    const [isLogin, setLogin] = useState(false);
    const inputTable = {
        loginUsername: setAccount,
        loginPassword: setPasswd,
    };

    const onInputChange = (e) => {
        const name = e.target.name;
        const val = e.target.value;

        inputTable[name](val);
    };

    function handleLogin() {
        apiHelper
            .login({ account, passwd })
            .then(() => apiHelper.profile())
            .then(() => setLogin(true))
            .catch(console.error);
    }

    return (
        <div>
            {!isLogin ? (
                <div id="formLogin" className="formContent show">
                    <div className="formContainer">
                        <div className="formField dynamic">
                            <input
                                type="text"
                                name="loginUsername"
                                id="loginUsername"
                                placeholder="Username"
                                autoComplete="off"
                                onChange={onInputChange}
                            />
                            <label htmlFor="loginUsername">Email</label>
                        </div>
                        <div className="formField dynamic">
                            <input
                                type="password"
                                name="loginPassword"
                                id="loginPassword"
                                placeholder="Password"
                                autoComplete="off"
                                onChange={onInputChange}
                            />
                            <label htmlFor="loginPassword">Password</label>
                            <div className="btnEye"></div>
                        </div>
                    </div>
                    <div className="btnContainer">
                        <button
                            onClick={handleLogin}
                            className="btnStyle sizeL green"
                        >
                            Login
                        </button>
                        <a href="#">Forgot password?</a>
                    </div>
                </div>
            ) : (
                <Redirect to={{ pathname: "/dashboard" }}></Redirect>
            )}
        </div>
    );
};

const Register = (props) => {
    const [email, setEmail] = useState("");
    const [passwd, setPasswd] = useState("");
    const [username, setUsername] = useState("");
    const [showPopup, setPopup] = useState(false);
    const inputTable = {
        registerEmail: setEmail,
        registerPassword: setPasswd,
        registerUsername: setUsername,
    };

    const onInputChange = (e) => {
        const name = e.target.name;
        const val = e.target.value;

        inputTable[name](val);
    };

    function handleRegister() {
        apiHelper
            .register({ email, name: username, passwd })
            .then(()=>{
                props.switchToLoginTab;
                setPopup(true);
            })
            .catch(console.error);
    }

    const clostPopup = () => {
        setPopup(false);
        props.switchToLoginTab();
    }

    return (
        <div id="formRegister" className="formContent show">
            <div className="formContainer">
                <div className="formField dynamic">
                    <input
                        type="text"
                        name="registerUsername"
                        id="registerUsername"
                        placeholder="Username"
                        onChange={onInputChange}
                    />
                    <label htmlFor="registerUsername">Username</label>
                </div>
                <div className="formField dynamic">
                    <input
                        type="text"
                        name="registerEmail"
                        id="registerEmail"
                        placeholder="Email"
                        onChange={onInputChange}
                    />
                    <label htmlFor="registerEmail">Email</label>
                </div>
                <div className="formField dynamic">
                    <input
                        type="password"
                        name="registerPassword"
                        id="registerPassword"
                        placeholder="Password"
                        onChange={onInputChange}
                    />
                    <label htmlFor="registerPassword">Password</label>
                    <div className="btnEye"></div>
                </div>
            </div>
            <div className="btnContainer">
                <button
                    onClick={handleRegister}
                    className="btnStyle sizeL green"
                >
                    Register
                </button>
            </div>
            {
                showPopup ? <FinishPopup clostPopup={clostPopup}></FinishPopup> : ''
            }
        </div>
    );
};

const Entry = () => {
    const [tabIndex, setTabIndex] = useState(0);

    const onTablChange = (event) => {
        if (event.target.id == "tabLogin") {
            setTabIndex(0);
        } else {
            setTabIndex(1);
        }
    };

    const switchToLoginTab = () => {
        setTabIndex(0);
    }

    return (
        <div id="entryPageWarpper">
            <div id="logoContainer">
                <img src={icLogoFull} title="Video Processing" />
            </div>
            <div id="loginContentContainer">
                <div id="tabContainer">
                    <div className="radioField">
                        <input
                            onChange={onTablChange}
                            type="radio"
                            name="contentTab"
                            id="tabLogin"
                            checked={tabIndex == 0}
                        />
                        <label htmlFor="tabLogin">Login</label>
                    </div>
                    <div className="radioField">
                        <input
                            onChange={onTablChange}
                            type="radio"
                            name="contentTab"
                            checked={tabIndex == 1}
                            id="tabRegister"
                        />
                        <label htmlFor="tabRegister">Register</label>
                    </div>
                </div>
                {tabIndex == 0 ? <Login /> : <Register switchToLoginTab={switchToLoginTab} />}
            </div>
        </div>
    );
};

export default Entry;
