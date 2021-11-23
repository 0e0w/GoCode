/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211026
*/
//4.逐字扫描
package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	for _, word := range words {
		fmt.Println(word)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("all time: %s\n", delta)
}
