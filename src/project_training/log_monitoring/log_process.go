package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	TimeLocal                    time.Time
	BytesSent                    int
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime    float64
}

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan *Message)
}

type LogProcess struct {
	rc     chan []byte   // read channel
	wc     chan *Message // write channel
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

func (w WriteToInfluxDB) Write(wc chan *Message) {
	for v := range wc {
		fmt.Println(v)
	}

}

func (l *LogProcess) Process() {
	// 172.0.0.12 - - [04/Mar/2018:13:49:53 +0000] http "GET /foo?query=t HTTP/1.0" 200 2133 "-" "KeepAliveClient" "-" 1.005 1.854
	// ([\d.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)
	r := regexp.MustCompile(`([\d.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^]]+)]\s+([a-z]+)\s+"([^"]+)"\s+(\d{3})\s+(\d+)\s+"([^"]+)"\s+"(.*?)"\s+"([\d.-]+)"\s+([\d.-]+)\s+([\d.-]+)`)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	for v := range l.rc {
		ret := r.FindStringSubmatch(string(v))
		if len(ret) != 14 {
			log.Fatalln("FindStringSubmatch fail: ", string(v))
			continue
		}

		message := &Message{}
		t, err := time.ParseInLocation("02/Jan/2006:15:04:05 +0000", ret[4], loc)
		if err != nil {
			log.Fatalln("ParseInLocation fail: ", err.Error(), ret[4])
		}

		message.TimeLocal = t
		byteSent, _ := strconv.Atoi(ret[8])
		message.BytesSent = byteSent

		// Get /foo?query=t HTTP/1.0
		reqSli := strings.Split(ret[6], " ")
		if len(reqSli) != 3 {
			log.Fatalln("strings.Split fail", ret[6])
			continue
		}

		url, err := url.Parse(reqSli[1])
		if err != nil {
			log.Fatalln("url parse fail:", err)
			continue
		}

		message.Path = url.Path
		message.Scheme = ret[5]
		message.Status = ret[7]

		upstreamTime, _ := strconv.ParseFloat(ret[12], 64)
		requestTime, _ := strconv.ParseFloat(ret[13], 64)
		message.UpstreamTime = upstreamTime
		message.RequestTime = requestTime
		l.wc <- message
	}

}

func main() {
	reader := &ReadFromFile{path: "./src/project_training/log_monitoring/access.log"}
	writer := &WriteToInfluxDB{dbSn: "username&password"}

	lp := &LogProcess{
		rc:     make(chan []byte),
		wc:     make(chan *Message),
		reader: reader,
		writer: writer,
	}

	go lp.reader.Read(lp.rc)
	go lp.Process()
	go lp.writer.Write(lp.wc)

	time.Sleep(60 * time.Second)
}
