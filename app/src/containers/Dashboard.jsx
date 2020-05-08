import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import icLogoIcon from "assets/images/ic_logo_icon.svg";
import icUser from "assets/images/ic_user.svg";
import VideoList from "../containers/VedioList";
import VideoTransfer from "../containers/VideoTransfer";
import apiHelper from "../services/api";

const NavMenu = (props) => {
    const [tabSelected, setSeleted] = useState(0);
    const switchToTransfer = () => {
        setSeleted(1);
        props.switchTab(1);
    };
    const switchToList = () => {
        setSeleted(0);
        props.switchTab(0);
    };
    return (
        <div>
            <nav>
                <div id="navLogoContainer">
                    <a href="https://www.vysioneer.com">
                        <img src={icLogoIcon} title="VYSIONEER" />
                    </a>
                </div>
                <ul>
                    <li
                        id="navUpload"
                        className={tabSelected == 1 ? "select" : ""}
                        onClick={switchToTransfer}
                    ></li>
                    <li
                        id="navList"
                        className={tabSelected == 0 ? "select" : ""}
                        onClick={switchToList}
                    ></li>
                </ul>
                <div id="navUserContainer">
                    <img
                        src={icUser}
                        onClick={() => {
                            apiHelper.logout().then(() => {
                                window.location.reload(false);
                            });
                        }}
                    />
                </div>
            </nav>
        </div>
    );
};

export const Dashboard = (props) => {
    const [tabIdx, setTabIdx] = useState(0);

    const switchTab = (idx) => {
        console.log(idx);
        setTabIdx(idx);
    };

    return (
        <div className="dashboard">
            <div id="dashboardPageWarpper">
                <NavMenu switchTab={switchTab} />
                <div className="pageContentContainer" id="uploadPageContent">
                    <div className="topMessageContainer">
                        Hi, {props.user.name}!
                    </div>
                    {tabIdx === 0 ? (
                        <VideoList />
                    ) : (
                        <VideoTransfer
                            goToVideoList={() => {
                                switchTab(0);
                            }}
                        />
                    )}
                </div>
            </div>
        </div>
    );
};
