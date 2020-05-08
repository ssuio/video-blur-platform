package job

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"vysioneer-assignment/model"
)

var jobManager *TaskManager

// TaskManager handle task seq and limit
type TaskManager struct {
	number int
	list   []VideoTask
}

type VideoTaskHandler func(videoID string) (interface{}, error)
type TaskStatus int

func (tm *TaskManager) AddVideoJob(video model.Video, handler VideoTaskHandler) error {
	task := VideoTask{
		video.ID,
		nil,
		nil,
		0,
		video,
		handler,
	}
	fmt.Println("Add job")
	tm.list = append(tm.list, task)
	fmt.Println(tm.list)

	return nil
}

func (tm *TaskManager) AddFaceBlurVideoJob(video model.Video) error {
	return tm.AddVideoJob(video, FaceBlurHandler)
}

func (tm *TaskManager) doVideoJob() {
	fmt.Println(tm.list)
	if len(tm.list) > 0 {
		fmt.Println("doing job")
		vj := tm.list[0]
		vj.Status = 1
		vj.video.Status = "processing"
		dm := model.GetDataManager()
		dm.UpdateVideo(vj.video)
		vj.Result, vj.Error = vj.Handler(vj.ID)
		if vj.Error != nil {
			vj.video.Status = "error"
		} else {
			vj.video.Status = "done"
		}
		dm.UpdateVideo(vj.video)
		vj.Status = 2
		tm.list = tm.list[1:]
	}
}

func (tm *TaskManager) CheckVideoJob(videoID string) int {
	for idx, v := range tm.list {
		if v.ID == videoID {
			return idx + 1
		}
	}

	return -1
}

func (tm *TaskManager) Run() {
	fmt.Println("Job manager running...")
	go func() {
		for {
			time.Sleep(10 * time.Second)
			fmt.Println("Check Job")
			tm.doVideoJob()
		}
	}()
}

type VideoTask struct {
	ID      string
	Result  interface{}
	Error   error
	Status  TaskStatus // 0=pending, 1=processing, 2=done, 3=failed
	video   model.Video
	Handler VideoTaskHandler
}

func FaceBlurHandler(videoID string) (interface{}, error) {
	fmt.Println("face blur start")
	source := "/data/tmp/" + videoID + ".mp4"
	dist := "/data/videos/" + videoID + ".mp4"
	out, err := exec.Command(
		"docker", "run",
		"-v", os.Getenv("HOST_DATA_DIR")+":/data",
		"-i", "face-blur",
		"-i", source,
		"-o", dist).Output()
	if err != nil {
		fmt.Println(err)
		fmt.Printf("FaceBlur result %s\n", out)
		return nil, err
	}
	return nil, nil
}

func GetJobManager() *TaskManager {
	if jobManager == nil {
		jobManager = &TaskManager{
			1,
			[]VideoTask{},
		}
	}
	return jobManager
}
