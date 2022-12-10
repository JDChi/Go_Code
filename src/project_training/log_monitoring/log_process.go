package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan string)
}

type LogProcess struct {
	rc     chan []byte // read channel
	wc     chan string // write channel
	reader Reader
	writer Writer
}

type ReadFromFile struct {
	path string // read file path
}

func (r ReadFromFile) Read(rc chan []byte) {

	// open file
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s ", err.Error()))
	}

	// seek to the end of the file, 2 which in whence means relative to the end
	f.Seek(0, 2)
	rd := bufio.NewReader(f)

	// read line by line
	for {
		// reads until the first occurrence of delim in the input, here is \n
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
		}
		// ignore \n in the end
		rc <- line[:len(line)-1]
	}

}

type WriteToInfluxDB struct {
	dbSn string // database data source
}

func (w WriteToInfluxDB) Write(wc chan string) {
	for v := range wc {
		fmt.Println(v)
	}

}

func (l *LogProcess) Process() {
	for v := range l.rc {
		l.wc <- strings.ToUpper(string(v))
	}

}

func main() {
	reader := &ReadFromFile{path: "./src/project_training/log_monitoring/access.log"}
	writer := &WriteToInfluxDB{dbSn: "username&password"}

	lp := &LogProcess{
		rc:     make(chan []byte),
		wc:     make(chan string),
		reader: reader,
		writer: writer,
	}

	go lp.reader.Read(lp.rc)
	go lp.Process()
	go lp.writer.Write(lp.wc)

	time.Sleep(30 * time.Second)
}
