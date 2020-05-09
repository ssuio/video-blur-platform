import React, { useState, useRef } from "react";
import apiHelper from "../services/api";

const Step1 = (props) => {
    const [name, setName] = useState("");
    const [desc, setDesc] = useState("");
    const [file, setFiles] = useState(null);
    const inputRef = useRef();
    const inputTable = {
        uploadFileName: setName,
        uploadNotes: setDesc,
    };

    const next = () => {
        console.log("step1 next page");
        props.setup({
            name,
            desc,
            file,
        });
        // apiHelper.transferVideo({
        //   name,
        //   desc,
        //   file: inputRef.current.files[0]
        // })
        props.handleSteps(1);
    };

    const onInputChange = (e) => {
        const name = e.target.name;
        const val = e.target.value;

        inputTable[name](val);
    };

    const fileUploadButton = () => {
        document.getElementById("fileButton").click();
        // document.getElementById("fileButton").onchange = () => {
        //     setFile(document.getElementById("fileButton").value);
        // };
    };

    const handleFileChange = (e) => {
        console.log(e.target.input);
        apiHelper.transferVideo({
            name,
            desc,
            file: e.target.input.files[0],
        });
        props.setup({
            file: e.target.files[0],
        });
    };

    return (
        <div>
            <div id="uploadZoneContainer">
                <div id="uploadZone">
                    <div id="uploadZoneMessage">
                        <div id="uploadZoneIcon"></div>
                        { file ? file.name : "Drag and Drop to upload" }
                    </div>
                    <input
                        id="fileButton"
                        type="file"
                        onChange={() => setFiles(inputRef.current.files[0])}
                        ref={inputRef}
                        hidden
                    />
                    <div id="uploadZoneBtn">
                        <button
                            className="btnStyle sizeS white"
                            onClick={fileUploadButton}
                        >
                            Select
                        </button>
                    </div>
                    <div id="uploadAnnotation">MP4 supported</div>
                </div>
            </div>
            <div id="uploadFormContainer">
                <div id="uploadFileNameContainer">
                    <div className="formField fixed">
                        <input
                            type="text"
                            name="uploadFileName"
                            id="uploadFileName"
                            placeholder="Please type your file name"
                            onChange={onInputChange}
                        />
                        <label htmlFor="uploadFileName">File Name</label>
                    </div>
                    <div id="uploadFileType">.mp4</div>
                </div>
                <div id="uploadNotesTextareaContainer">
                    <textarea
                        name="uploadNotes"
                        id="uploadNotes"
                        rows="3"
                        onChange={onInputChange}
                    />
                    <label htmlFor="uploadNotes">Notes</label>
                </div>
            </div>
            <div className="processBtnContainer">
                <button
                    className="btnStyle sizeL green"
                    id="btnStepUpload"
                    onClick={next}
                >
                    Next
                </button>
            </div>
        </div>
    );
};

const Step2 = (props) => {
    const [perm, setPerm] = useState(false);
    const next = () => {
        props.handleSteps(2);
        props.setup({ perm });
    };

    const permOnChange = (e) => {
        console.log(e.target.checked);
        setPerm(e.target.checked);
    };
    return (
        <div>
            <div id="transferFileName">Brain_Tumor_01.mp4</div>
            <div id="transferFormContainer">
                <div className="transferFormRow">
                    <div className="transferFormTitle">Add Effect</div>
                    <div id="addEffectRadioContainer">
                        <div className="formRadioField">
                            <input
                                type="radio"
                                name="transferAddEffect"
                                id="transferAddEffectY"
                                defaultChecked
                            />
                            <label htmlFor="transferAddEffectY">
                                <div className="radioIcon">
                                    <div className="radioInsideCircle"></div>
                                </div>
                                <div className="radioText">Yes</div>
                            </label>
                        </div>
                        <div className="formRadioField">
                            <input
                                type="radio"
                                name="transferAddEffect"
                                id="transferAddEffectN"
                                disabled
                            />
                            <label htmlFor="transferAddEffectN">
                                <div className="radioIcon">
                                    <div className="radioInsideCircle"></div>
                                </div>
                                <div className="radioText">No</div>
                            </label>
                        </div>
                    </div>
                </div>
                <div className="transferFormRow" id="transferTypeForm">
                    <div className="transferFormTitle">Type</div>
                    <div id="typeCheckboxContainer">
                        <div className="formCheckboxField">
                            <input
                                type="checkbox"
                                name="transferType"
                                id="transferTypeMosaic"
                                defaultChecked
                            />
                            <label htmlFor="transferTypeMosaic">
                                <div className="checkboxIcon"></div>
                                <div className="checkboxText">Face Blur</div>
                            </label>
                        </div>
                        <div className="formCheckboxField">
                            <input
                                type="checkbox"
                                name="transferType"
                                id="transferTypeOther1"
                            />
                            <label htmlFor="transferTypeOther1">
                                <div className="checkboxIcon"></div>
                                <div className="checkboxText">Other1</div>
                            </label>
                        </div>
                        <div className="formCheckboxField">
                            <input
                                type="checkbox"
                                name="transferType"
                                id="transferTypeOther2"
                            />
                            <label htmlFor="transferTypeOther2">
                                <div className="checkboxIcon"></div>
                                <div className="checkboxText">Other2</div>
                            </label>
                        </div>
                    </div>
                </div>
                <div className="transferFormRow">
                    <div id="publicSwitchContainer">
                        <div className="transferFormTitle">Public</div>
                        <div className="switchField">
                            <label>
                                <input
                                    type="checkbox"
                                    onChange={permOnChange}
                                    checked={perm}
                                />
                                <span className="slider"></span>
                            </label>
                        </div>
                    </div>
                </div>
            </div>
            <div className="processBtnContainer">
                <button
                    className="btnStyle sizeL green"
                    id="btnStepTransfer"
                    onClick={next}
                >
                    Next
                </button>
            </div>
        </div>
    );
};

const Step3 = (props) => {
    const goToVideoList = () => {
        props.goToVideoList();
    };

    let params = props.getParams();
    console.log(params);
    apiHelper
        .transferVideo(params.name, params.desc, params.file, params.perm)
        .then(console.log)
        .catch(console.error);

    return (
        <div>
            <div id="completeMarkContainer">
                <div id="completeIconContainer"></div>
                Complete!
            </div>
            <div id="completeTextContainer">
                Your upload process is completed!
                <br />
                You may need to wait for awhile dealing the file,
                <br />
                Please check the status in the list table.
            </div>
            <div className="processBtnContainer">
                <button
                    className="btnStyle sizeL green"
                    id="btnStepComplete"
                    onClick={goToVideoList}
                >
                    Go to List
                </button>
            </div>
        </div>
    );
};

const VideoTransfer = (props) => {
    const [params, setParams] = useState({});
    const [step, setStep] = useState(0);
    const [wizardState, setWizardState] = useState({
        step: 1,
        upload: "now",
        transfer: "",
        complete: "",
    });

    const setup = (data) => {
        setParams({ ...params, ...data });
    };

    const getParams = (data) => {
        return params;
    };

    const handleSteps = (s) => {
        setStep(s);
        switch (s) {
            case 0:
                setWizardState({
                    step: 1,
                    upload: "now",
                    transfer: "",
                    complete: "",
                });
                return;
            case 1:
                setWizardState({
                    upload: "done",
                    transfer: "now",
                    complete: "",
                });
                return;
            case 2:
                setWizardState({
                    upload: "done",
                    transfer: "done",
                    complete: "now",
                });
                return;
        }
    };

    return (
        <div>
            <div className="dashboardContentContainer">
                <div id="progressBarContainer">
                    <ul>
                        <li
                            id="progressBarUpload"
                            className={wizardState.upload}
                        >
                            1<div className="progressText">Upload</div>
                        </li>
                        <li
                            id="progressBarTransfer"
                            className={wizardState.transfer}
                        >
                            2<div className="progressText">Transfer</div>
                        </li>
                        <li
                            id="progressBarComplete"
                            className={wizardState.complete}
                        >
                            3<div className="progressText">Complete</div>
                        </li>
                    </ul>
                </div>
                <div id="stepUploadContainer" className="stepContainer show">
                    {step === 0 ? (
                        <Step1 handleSteps={handleSteps} setup={setup} />
                    ) : (
                        ""
                    )}
                    {step === 1 ? (
                        <Step2 handleSteps={handleSteps} setup={setup} />
                    ) : (
                        ""
                    )}
                    {step === 2 ? (
                        <Step3
                            handleSteps={handleSteps}
                            goToVideoList={props.goToVideoList}
                            setup={setup}
                            getParams={getParams}
                        />
                    ) : (
                        ""
                    )}
                </div>
            </div>
        </div>
    );
};

export default VideoTransfer;
