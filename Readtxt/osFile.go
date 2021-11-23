/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211123
*/

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	filename := "urls.txt"
	fl, err := os.Open(filename)
	if err != nil {
		fmt.Println(filename, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("all time: %s\n", delta)
}
