/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211026
*/
//读文件方式三：先从文件读取到file, 在从file读取到Reader中，从Reader读取到buf, buf最终追加到[]byte
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	Read2()

}
//先从文件读取到file, 在从file读取到Reader中，从Reader读取到buf, buf最终追加到[]byte，这个排第三
func Read2() (string) {
	filenanme := "urls.txt"
	fi, err := os.Open(filenanme)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	r := bufio.NewReader(fi)
	var chunks []byte

	buf := make([]byte, 1024)

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		//fmt.Println(string(buf))
		chunks = append(chunks, buf...)
	}
	fmt.Println(string(chunks))
	return string(chunks)
}

