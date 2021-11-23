/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211026
*/
//2.以块的形式读取文件
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const BufferSize02 = 1000000

func main() {
	start := time.Now()

	filenanme := "urls.txt"
	file, err := os.Open(filenanme)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	buffer := make([]byte, BufferSize02)
	for {
		bytesread, err := file.Read(buffer)
		//_, err = file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		//fmt.Println("bytes read: ", bytesread)
		fmt.Println(string(buffer[:bytesread]))
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("all time: %s\n", delta)
}
