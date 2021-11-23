/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211123
*/

package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	start := time.Now()
	dstxt()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("all time: %s\n", delta)

}
func dstxt() {
	filenanme := "urls.txt"
	dstxt, err := ioutil.ReadFile(filenanme)
	if err != nil {
		fmt.Println(err)
	}
	dstxtstr := string(dstxt)
	fmt.Println(dstxtstr)
}
