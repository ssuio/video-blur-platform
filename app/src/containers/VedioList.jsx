import React, { useState, useEffect } from "react";
import apiHelper from "../services/api";


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
    const [perm, setPerm] = useState(props.video.Perm)
    const [openSubContent, setOpenSubContent] = useState(false)
    const [isEditable, setEditable] = useState(false)
    const [isVideoCheck, setVideoCheck] = useState(props.videoCheckList[props.videoIdx])
    const videoIdx = props.videoIdx

    const toDate = (date) => {
        date = new Date(date);
        let mm = date.getMonth() + 1; // getMonth() is zero-based
        let dd = date.getDate();
        return [(mm > 9 ? "" : "0") + mm, (dd > 9 ? "" : "0") + dd].join("-");
    };

    const handleVideoCheck = (e) => {
        let checked = e.target.checked;
        console.log(`set video ${videoIdx} check ${checked}`)
        setVideoCheck(checked)
        props.videoCheckList[videoIdx] = checked;
    }

    const handleSharelink = (e) => {
        let perm = e.target.checked;
        apiHelper.updateVideo(props.video.id, { perm })
            .then(() => {
                console.log(`set ${perm}`);
                setPerm(perm);
            })
    }

    const toggleSubContent = () => {
        let val = !openSubContent;
        setOpenSubContent(val);
    }

    const toggleEditable = () => {
        let val = !isEditable;
        setEditable(val);
    }

    const copyToClipboard = () => {
        let textField = document.createElement('textarea')
        textField.innerText = `https://localhost:9000/sharelink/${props.video.id}`
        document.body.appendChild(textField)
        textField.select()
        document.execCommand('copy')
        textField.remove()
    }

    return (
        <div className="listRow">
            <div className="mainContent">
                <div className="listCell cellCheckbox">
                    <div className="checkboxField">
                        <input
                            type="checkbox"
                            name={'tableFileList' + videoIdx}
                            id={'tableListRow1' + videoIdx}
                            // onChange={handleVideoCheck}
                            // checked={isVideoCheck}
                        />
                        <label htmlFor="tableListRow1"></label>
                    </div>
                </div>
                <div className="listCell cellDate">
                    {toDate(props.video.CreatedTime)}
                </div>
                <div className="listCell cellFileName">
                    <input type="text" defaultValue={props.video.name} disabled={!isEditable} />
                </div>
                <div className="listCell cellPublic">
                    <div className="switchField">
                        <label>
                            <input type="checkbox" checked={perm} onChange={handleSharelink} />
                            <span className="slider"></span>
                        </label>
                    </div>
                </div>
                <div className="listCell cellSize">{props.video.Size}</div>
                <div className={"listCell cellStatus " + props.video.Status}>{props.video.Status}</div>
                <div className="listCell cellEdit" onClick={toggleEditable}>
                    <div className="editIconContainer"></div>
                </div>
                <div className="listCell cellBtn" onClick={toggleSubContent}>
                    <div className={`cellBtnContainer ${openSubContent ? 'open' : ''}`}></div>
                </div>
            </div>
            <div className={`subContent ${openSubContent ? 'show' : ''}`}>
                <div className="subListRow rowNotes">
                    <div className="subListCell subListTitle">Notes</div>
                    <div className="subListCell notesTextarea">
                        <textarea
                            name="tableNotes1"
                            id="tableNotes1"
                            rows="2"
                            defaultValue={props.video.description}
                            disabled={!isEditable}
                        />
                    </div>
                </div>
                <div className="subListRow rowLink">
                    <div className="subListCell subListTitle">Link</div>
                    <div className="subListCell linkInput" onClick={copyToClipboard}>
                        <input
                            type="text"
                            defaultValue={`https://localhost:9000/sharelink/${props.video.id}`}
                            disabled
                        />
                        <div className="linkCopyBtn"></div>
                    </div>
                </div>
            </div>
        </div>
    );
};

const VedioList = (props) => {
    const [videoList, setVideoList] = useState([]);
    const [videoCheckList, setVideoCheckList] = useState([]);

    const refreshVideoList = () => {
        apiHelper.videos().then((videos) => {
            setVideoCheckList(videos.map(() => false));
            setVideoList(videos);
        });
    };

    const deleteVideo = async () => {
        console.log(videoCheckList);
        for (let idx in videoCheckList) {
            console.log(`check ${idx}`);
            if (videoCheckList[idx]) {
                console.log("delete");
                console.log(videoList);
                console.log(idx);
                // await apiHelper.deleteVideo(videoList[v])
            }
        }
    };

    useEffect(refreshVideoList, []);

    return (
        <div>
            <div className="dashboardContentContainer">
                <header>File List</header>
                <div id="listTableBtnContainer">
                    <button
                        className="btnStyle sizeS whiteLine"
                        onClick={refreshVideoList}
                    >
                        Refresh
                    </button>
                    <button
                        className="btnStyle sizeS redLine"
                        onClick={deleteVideo}
                    >
                        Delete
                    </button>
                </div>
                <div id="listTable">
                    <div className="listRow title">
                        <div className="listCell cellCheckbox"> </div>
                        <div className="listCell cellDate">Date</div>
                        <div className="listCell cellFileName">File Name</div>
                        <div className="listCell cellPublic">Public</div>
                        <div className="listCell cellSize">Size</div>
                        <div className="listCell cellStatus">Status</div>
                        <div className="listCell cellEdit"> </div>
                        <div className="listCell cellBtn"> </div>
                    </div>
                    {videoList.map((v, idx) => {
                        console.log("map " + idx);
                        return (
                            <ListRow
                                key={idx}
                                video={v}
                                videoCheckList={videoCheckList}
                                videoIdx={idx}
                            />
                        );
                    })}
                </div>
            </div>
        </div>
    );
};

export default VedioList;