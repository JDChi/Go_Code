package main

import (
	"fmt"
	"strings"
	"time"
)

type Reader interface {
	Read(rc chan string)
}

type Writer interface {
	Write(wc chan string)
}

type LogProcess struct {
	rc     chan string // read channel
	wc     chan string // write channel
	reader Reader
	writer Writer
}

type ReadFromFile struct {
	path string // read file path
}

func (r ReadFromFile) Read(rc chan string) {
	line := "message"
	rc <- line
}

type WriteToInfluxDB struct {
	dbSn string // database data source
}

func (w WriteToInfluxDB) Write(wc chan string) {
	fmt.Println(<-wc)
}

func (l *LogProcess) Process() {
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

func main() {
	reader := &ReadFromFile{path: "/tmp/access.log"}
	writer := &WriteToInfluxDB{dbSn: "username&password"}

	lp := &LogProcess{
		rc:     make(chan string),
		wc:     make(chan string),
		reader: reader,
		writer: writer,
	}

	go lp.reader.Read(lp.rc)
	go lp.Process()
	go lp.writer.Write(lp.wc)

	time.Sleep(2 * time.Second)
}
