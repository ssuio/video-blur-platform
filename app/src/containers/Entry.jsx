import React from "react";
import Tabs from "@material-ui/core/Tabs";
import Tab from "@material-ui/core/Tab";
import Box from '@material-ui/core/Box';
import PropTypes from 'prop-types';

import { Login, Register } from "./User"

const TabPanel = (props) => {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`scrollable-auto-tabpanel-${index}`}
      aria-labelledby={`scrollable-auto-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box p={3}>
          {children}
        </Box>
      )}
    </div>
  );
}

TabPanel.propTypes = {
  children: PropTypes.node,
  index: PropTypes.any.isRequired,
  value: PropTypes.any.isRequired,
};

function a11yProps(index) {
  return {
    id: `scrollable-auto-tab-${index}`,
    'aria-controls': `scrollable-auto-tabpanel-${index}`,
  };
}

const Entry = () => {
  const [tabIndex, setTabIndex] = React.useState(0);

  const handleTabChange = (event, newIndex) => {
    setTabIndex(newIndex);
  };

  const goToLoginTab = () =>{
    setTabIndex(0);
  }

  return <div>
    <div style={{
      position:'relative',
      left:'33%',
      top:'15vh',
      width:'33%',
      backgroundColor: 'rgb(19, 21, 26, 0.3)',
      padding:'15px'
    }}>
      <Tabs
        value={tabIndex}
        onChange={handleTabChange}
        indicatorColor="primary"
        textColor="primary"
        variant="scrollable"
        scrollButtons="auto"
      >
        <Tab label='Login' {...a11yProps(0)} />
        <Tab label='Register' {...a11yProps(1)} />

      </Tabs>
      <br /><br />
      <TabPanel value={tabIndex} index={0}>
        <div>
          <Login />
        </div>
      </TabPanel>
      <TabPanel value={tabIndex} index={1}>
        <div>
          <Register 
            goToLoginTab = {goToLoginTab}
          />
        </div>
      </TabPanel>

    </div>
  </div>
}

export default Entry