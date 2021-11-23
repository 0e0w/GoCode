/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211026
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const BufferSizeBufferSize = 100

type chunkchunk struct {
	bufsize int
	offset  int64
}

func main() {
	start := time.Now()
	filenanme := "urls.txt"
	file, err := os.Open(filenanme)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// Returns a boolean based on whether there's a next instance of `\n`
	// character in the IO stream. This step also advances the internal pointer
	// to the next position (after '\n') if it did find that token.
	for {
		read := scanner.Scan()
		if !read {
			break
		}
		//fmt.Println("read byte array: ", scanner.Bytes())
		fmt.Println( scanner.Text())
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("all time: %s\n", delta)
}
