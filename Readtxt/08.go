/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211026
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filenanme := "urls1.txt"
	file, err := os.Open(filenanme)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Println(string(content))
}
