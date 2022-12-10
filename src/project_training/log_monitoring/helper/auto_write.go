package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	filePath := "./src/project_training/log_monitoring/access.log"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
	}

	defer file.Close()

	write := bufio.NewWriter(file)
	for i := 0; i < 20; i++ {
		time.Sleep(2 * time.Second)
		write.WriteString("172.0.0.12 - - [04/Mar/2018:13:49:53 +0000] http \"GET /foo?query=t HTTP/1.0\" 200 2133 \"-\" \"KeepAliveClient\" \"-\" 1.005 1.854\n")
		write.Flush()
	}

}
