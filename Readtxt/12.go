/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211026
*/
//读文件方式二：先从文件读取到file中，在从file读取到buf, buf在追加到最终的[]byte
package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	Read1()
}

func Read1()  (string){
	//获得一个file
	filenanme := "urls.txt"
	f, err := os.Open(filenanme)
	if err != nil {
		fmt.Println("read fail")
		return ""
	}

	//把file读取到缓冲区中
	defer f.Close()
	var chunk []byte
	buf := make([]byte, 1024)

	for {
		//从file读取到buf中
		n, err := f.Read(buf)
		if err != nil && err != io.EOF{
			fmt.Println("read buf fail", err)
			return ""
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}
	return string(chunk)
	//fmt.Println(string(chunk))
}

