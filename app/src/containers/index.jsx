import React from "react";
import { hot } from "react-hot-loader";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import VedioView from "./VedioPage";

const Main = () => <h1>Main</h1>;

const NavMenu = () => (
  <div>
    <ul>
      <li>
        <Link to="/">Main</Link>
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
    path: "/user",
    component: () => (
      <div>
        <h1>User</h1>
      </div>
    ),
  },
  {
    key: 1,
    path: "/video",
    component: VedioView,
  },
];

function App() {
  return (
    <div>
      <div className="App">
        <Router>
          <NavMenu />
          <Switch>
            <Route exact path="/" component={Main} />
            {routes.map((route) => (
              <Route {...route} />
            ))}
          </Switch>
        </Router>
      </div>
    </div>
  );
}

export default hot(module)(App)