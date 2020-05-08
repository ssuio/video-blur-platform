import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import icLogoIcon from "assets/images/ic_logo_icon.svg";
import icUser from "assets/images/ic_user.svg";
import VideoList from "../containers/VedioList";
import VideoTransfer from "../containers/VideoTransfer";

const NavMenu = () => {
    return (
        <div>
            <nav>
                <div id="navLogoContainer">
                    <Link to="https://www.vysioneer.com">
                        <img src={icLogoIcon} title="VYSIONEER" />
                    </Link>
                </div>
                <ul>
                    <li id="navUpload"></li>
                    <li id="navList" className="select"></li>
                </ul>
                <div id="navUserContainer">
                    <img src={icUser} />
                </div>
            </nav>
        </div>
    );
};

export const Dashboard = (props) => {
    const [tabIdx, setTabIdx] = useState(1);

    const switchTab = (idx) => {
        setTabIdx(idx);
    };

    const renderTab = () => {
        console.log("render " + tabIdx);
        switch (tabIdx) {
            case 0:
                return <VideoList />;
            case 1:
                return (
                    <VideoTransfer
                        goToVideoList={() => {
                            switchTab(0);
                        }}
                    />
                );
        }
    };
    return (
        <div className="dashboard">
            <div id="dashboardPageWarpper">
                <NavMenu switchTab={switchTab} />
                <div className="pageContentContainer" id="uploadPageContent">
                    <div className="topMessageContainer">
                        Hi, {props.user.name}!
                    </div>
                    {renderTab()}
                </div>
            </div>
        </div>
    );
};
