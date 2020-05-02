package job

import (
	"os/exec"
)

// TaskManager handle task seq and limit 
type TaskManager struct {
	number int
	list []Task
}

func (tm TaskManager) addJob(){
	
}

type TaskStatus int

// Task a job
type Task struct {
	Result interface{}
	err error
	Status TaskStatus // 0=pending, 1=processing, 2=done, 3=failed
	Handler func
}

func (t Task)start(){
	t.Handler()
}

func FaceBlurHandler(videoDir, videoID, targetVideoID string) error{
	out, err := exec.Command(
			fmt.Sprintf("docker run -v %s:/tmp -i face-recongnition -i %s -o %s", 
			videoDir, "/tmp/"+videoID+".mp4", "/tmp/"+targetVideoID+".mp4")
		).Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("FaceBlur result %s\n", out)
	return nil
}