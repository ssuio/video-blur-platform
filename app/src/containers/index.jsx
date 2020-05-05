import React from "react";
import { hot } from "react-hot-loader";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  Redirect,
} from "react-router-dom";
import VedioView from "./Vedio";
import { Login, Register } from "./User";

const NavMenu = () => (
  <div>
    <ul>
      <li>
        <Link to="/login">Login</Link>
      </li>
      <li>
        <Link to="/user">Account</Link>
      </li>
      <li>
        <Link to="/video">Video</Link>
      </li>
    </ul>
    <hr />
  </div>
);

const routes = [
  {
    key: 0,
    path: "/login",
    component: Login,
  },
  {
    key: 1,
    path: "/user",
    component: () => (
      <div>
        <h1>User</h1>
      </div>
    ),
  },
  {
    key: 2,
    path: "/video",
    component: VedioView,
  },
];

const Page404 = () => <h1> Page 404 </h1>;
const Loading = () => <h1>Loading ... </h1>;

const PrivateRoute = ({ component: Component, ...rest }) => {
  const [loading, setLoading] = React.useState(true);
  const [auth, setAuth] = React.useState(false);

  const checkAuth = () => {
    axios.get("http://127.0.0.1:9000/user")
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

  checkAuth();

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
          <NavMenu />
          <Switch>
            <Route path="/login" component={Login} />
            <Route path="/register" component={Register} />
            {routes.map((route) => (
              <PrivateRoute {...route} />
            ))}
            <Route path="*" exact={true} component={Page404} />
          </Switch>
        </Router>
      </div>
    </div>
  );
}

export default hot(module)(App);
