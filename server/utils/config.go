package utils

import (
	"os"
)

func GetDataDir (){
	return os.Getenv("DATA_DIR")
}