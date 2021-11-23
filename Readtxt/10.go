/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211026
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filenanme := "urls.txt"
	f, err := os.Open(filenanme)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	chunks := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Println(string(chunks))
}
