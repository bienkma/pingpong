package pingpong

import (
	"os"
	"log"
)

func Log(err interface{}){
	f, err_opening := os.OpenFile("/var/log/pingpong.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
	if err_opening != nil{
		panic(err_opening)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Printf("Error: %v", err)
}