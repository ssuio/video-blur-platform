import React, { useState, useEffect } from "react";
import {
    Link,
  } from "react-router-dom";
import {Logout} from "./User"

const NavMenu = () => (
    <div>
      <ul>
        <li>
          <Link to="/user">Account</Link>
        </li>
        <li>
          <Link to="/video">Video Transfer</Link>
        </li>
        <li>
          <Link to="/videoList">Video List</Link>
        </li>
      </ul>
      <hr />
    </div>
  );

const Dashboard = (props) => {
  return <div> 
    This is dashbaord page
    {props.children}  
  </div>
}

export default Dashboard