import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import icLogoIcon from "assets/images/ic_logo_icon.svg";
import icUser from "assets/images/ic_user.svg";
import apiHelper from "../services/api";

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

const ListRow = (props) => {
    // const [video, setVideo] = useState({});
    // const [status, setStatus] = useState("");

    // useEffect(() => {
    //     let tV = props.video;
    //     setVideo({
    //         id: tV.id,
    //         name: tV.name,
    //         description: tV.description,
    //         size: tV.Size,
    //         createdTime: tV.CreatedTime,
    //         imageUrl: tV.ImageUrl,
    //         ownerID: tV.OwnerID,
    //         perm: tV.perm,
    //         status: tV.Status,
    //     });
    //     if (props.video.id == "1") {
    //         console.log(`id 1 set status ${tV.Status}`);
    //     }
    //     setStatus(tV.Status);
    // }, []);

    const toDate = (date) => {
        date = new Date(date);
        let mm = date.getMonth() + 1; // getMonth() is zero-based
        let dd = date.getDate();
        return [(mm > 9 ? "" : "0") + mm, (dd > 9 ? "" : "0") + dd].join("-");
    };
    
    return (
        <div className="listRow">
            <div className="mainContent">
                <div className="listCell cellCheckbox">
                    <div className="checkboxField">
                        <input
                            type="checkbox"
                            name="tableFileList"
                            id="tableListRow1"
                        />
                        <label htmlFor="tableListRow1"></label>
                    </div>
                </div>
                <div className="listCell cellDate">
                    {toDate(props.video.CreatedTime)}
                </div>
                <div className="listCell cellFileName">
                    <input type="text" value={props.video.name} disabled />
                </div>
                <div className="listCell cellPublic">
                    <div className="switchField">
                        <label>
                            <input type="checkbox" />
                            <span className="slider"></span>
                        </label>
                    </div>
                </div>
                <div className="listCell cellSize">{props.video.Size}</div>
                <div className={"listCell cellStatus "+props.video.Status}>{props.video.Status}</div>
                <div className="listCell cellEdit">
                    <div className="editIconContainer"></div>
                </div>
                <div className="listCell cellBtn">
                    <div className="cellBtnContainer"></div>
                </div>
            </div>
            <div className="subContent">
                <div className="subListRow rowNotes">
                    <div className="subListCell subListTitle">Notes</div>
                    <div className="subListCell notesTextarea">
                        <textarea
                            name="tableNotes1"
                            id="tableNotes1"
                            rows="2"
                            disabled
                        >
                            {props.video.description}
                        </textarea>
                    </div>
                </div>
                <div className="subListRow rowLink">
                    <div className="subListCell subListTitle">Link</div>
                    <div className="subListCell linkInput">
                        <input
                            type="text"
                            value="https://www.vysioneer.com/Brain_Tumor_01.mp4"
                        />
                        <div className="linkCopyBtn"></div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export const Dashboard = (props) => {
    const [videoList, setVideoList] = useState([]);

    const refreshVideoList = () => {
        apiHelper.videos().then((videos) => {
            setVideoList(videos);
        });
    };

    useEffect(refreshVideoList, []);

    return (
        <div className="dashboard">
            <div id="dashboardPageWarpper">
                <NavMenu />
                <div className="pageContentContainer" id="uploadPageContent">
                    <div className="topMessageContainer">Hi, Noah!</div>
                    <div className="dashboardContentContainer">
                        <header>File List</header>
                        <div id="listTableBtnContainer">
                            <button
                                className="btnStyle sizeS whiteLine"
                                onClick={refreshVideoList}
                            >
                                Refresh
                            </button>
                            <button className="btnStyle sizeS redLine" disabled>
                                Delete
                            </button>
                        </div>
                        <div id="listTable">
                            <div className="listRow title">
                                <div className="listCell cellCheckbox"> </div>
                                <div className="listCell cellDate">Date</div>
                                <div className="listCell cellFileName">
                                    File Name
                                </div>
                                <div className="listCell cellPublic">
                                    Public
                                </div>
                                <div className="listCell cellSize">Size</div>
                                <div className="listCell cellStatus">
                                    Status
                                </div>
                                <div className="listCell cellEdit"> </div>
                                <div className="listCell cellBtn"> </div>
                            </div>
                            {videoList.map((v, idx) => {
                                return <ListRow key={idx} video={v} />;
                            })}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};
