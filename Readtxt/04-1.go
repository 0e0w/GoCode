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
)

func main() {
	filenanme := "urls.txt"
	file, err := os.Open(filenanme)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	// initial size of our wordlist
	bufferSize := 50
	words := make([]string, bufferSize)
	pos := 0
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			// This error is a non-EOF error. End the iteration if we encounter
			// an error
			fmt.Println(err)
			break
		}
		words[pos] = scanner.Text()
		pos++
		if pos >= len(words) {
			// expand the buffer by 100 again
			newbuf := make([]string, bufferSize)
			words = append(words, newbuf...)
		}
	}
	fmt.Println("word list:")
	// we are iterating only until the value of "pos" because our buffer size
	// might be more than the number of words because we increase the length by
	// a constant value. Or the scanner loop might've terminated due to an
	// error prematurely. In this case the "pos" contains the index of the last
	// successful update.
	for _, word := range words[:pos] {
		fmt.Println(word)
	}
}
