package main

import (
	"fmt"
	"log"
	"proto-course/basic"
	"time"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(&logWriter{})

	// basic.BaseUser()
	basic.BasicUserGroup()
}

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("2006-01-02 15:04:05") + " " + string(bytes))
}
