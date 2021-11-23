/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211026
*/
//读取文件的整个目录
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	filelist, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, fileinfo := range filelist {
		if fileinfo.Mode().IsRegular() {
			bytes, err := ioutil.ReadFile(fileinfo.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Bytes read: ", len(bytes))
			fmt.Println("String read: ", string(bytes))
		}
	}
}
