package job

import (
	"fmt"
	"os"
	"os/exec"
)

// TaskManager handle task seq and limit
type TaskManager struct {
	number int
	list   []Task
}

func (tm TaskManager) addJob() {

}

type TaskStatus int

// Task a job
type Task struct {
	Result  interface{}
	err     error
	Status  TaskStatus // 0=pending, 1=processing, 2=done, 3=failed
	Handler func(args ...interface{})
}

func (t Task) start() {
	t.Handler()
}

func FaceBlurHandler(videoID string) error {
	source := "/data/tmp/"+videoID+".mp4"
	dist := "/data/videos/"+videoID+".mp4"
	out, err := exec.Command(
		"docker", "run", 
		"-v", os.Getenv("DATA_DIR") + ":/data",
		"-i", "face-blur", 
		"-i", source, 
		"-o", dist).Output()
	if err != nil {
		fmt.Println(err)
		fmt.Printf("FaceBlur result %s\n", out)
		return err
	}
	return nil
}
