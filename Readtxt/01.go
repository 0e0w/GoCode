/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211026
*/
//1.读取整个文件到内存中
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	filenanme := "urls.txt"
	file, err := os.Open(filenanme)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(bytesread)
	fmt.Println(string(buffer))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("all time: %s\n", delta)
}
