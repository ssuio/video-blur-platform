import React, {useState, useEffect} from "react";
import apiHelper from "../services/api";

export const Login = () => {
    const [loading, setLoading] = useState(true);
    const [isLogin, setLogin] = useState(false);
    const [user, setUser] = useState({});
    const [account, setAccount] = useState("");
    const [passwd, setPasswd] = useState("");

    const checkLogin = () => {
        apiHelper
            .profile()
            .then((user) => {
                setLogin(true);
                setLoading(false);
                setUser(user);
                return;
            })
            .catch(() => {
                setLogin(false);
                setLoading(false);
            });
    };

    useEffect(()=>{
        checkLogin()
    })

    function login() {
        apiHelper
            .login(account, passwd)
            .then(() => apiHelper.profile())
            .then(user => {
                console.log(user)
                setUser(user)
            })
            .then(() => setLogin(true))
            .catch(console.error);
    }

    function logout() {
        apiHelper
            .logout()
            .then(() => {
                setUser({})
                setLogin(false);
            })
            .catch(console.error);
    }

    const accountOnChange = (e) => {
        setAccount(e.target.value);
    };

    const passwdOnChange = (e) => {
        setPasswd(e.target.value);
    };

    return (
        <div>
            {!loading && user.email ? (
                <div>
                    <label>{user.name}</label>
                    <br />
                    <label>{user.email}</label>
                    <br />
                    <button onClick={logout}>Logout</button>
                </div>
            ) : (
                <div>
                    <h1>Login</h1>
                    <br />
                    <label>Account</label>
                    <input
                        onChange={accountOnChange}
                        value={account}
                        type="text"
                    />
                    <br />
                    <label>Password</label>
                    <input
                        onChange={passwdOnChange}
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
export const Register = () => <h1> Regeister </h1>;
