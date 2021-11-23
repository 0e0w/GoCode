/*
0 error(s),0 warning(s)
Team:0e0w Security Team
Author:0e0wTeam[at]gmail.com
Datetime:20211026
*/
//读文件方式四：读取到file中，再利用ioutil将file直接读取到[]byte中

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
	Read3()
}

// Read3 读取到file中，再利用ioutil将file直接读取到[]byte中, 这是最优
func Read3() {
	filenanme := "urls.txt"
	f, err := os.Open(filenanme)
	if err != nil {
		fmt.Println("read file fail", err)
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
	}
	fmt.Println(string(fd))
}

