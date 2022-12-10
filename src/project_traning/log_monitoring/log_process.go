package main

import (
	"fmt"
	"strings"
	"time"
)

type LogProcess struct {
	path string      // read file path
	dbSn string      // database data source
	rc   chan string // read channel
	wc   chan string // write channel
}

func (l *LogProcess) ReadFromFile() {
	line := "message"
	l.rc <- line
}

func (l *LogProcess) Process() {
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

func (l *LogProcess) WriteToDB() {
	fmt.Println(<-l.wc)
}

func main() {
	lp := &LogProcess{
		rc:   make(chan string),
		wc:   make(chan string),
		path: "/tmp/access.log",
		dbSn: "username&password",
	}

	go lp.ReadFromFile()
	go lp.Process()
	go lp.WriteToDB()

	time.Sleep(2 * time.Second)
}
