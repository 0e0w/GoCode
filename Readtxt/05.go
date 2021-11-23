/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211026
*/

//5.将长字符串拆分为单词

package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	longstring := "This is a very long string. Not."
	var words []string
	scanner := bufio.NewScanner(strings.NewReader(longstring))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	fmt.Println("word list:")
	for _, word := range words {
		fmt.Println(word)
	}
}
