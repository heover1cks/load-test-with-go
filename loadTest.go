package main

import (
"fmt"
"github.com/codeskyblue/go-sh"
"strconv"
"time"
)

var mainUri = "LBADDRESS"
var subPath = "URIPATH"
var patchStr = "PARAMETER"

func curlExec(id int) {
	path := mainUri + subPath + strconv.Itoa(id) + patchStr
	fmt.Println(path)
	for {
		sh.Command("curl", "-X", "PATCH", path, "-H", "accept: */*").Run()
		time.Sleep(time.Millisecond * 100)
		fmt.Println("ID:" + strconv.Itoa(id) + " PATCHed")
	}
}

func main() {
	for i := 1; i < 2000; i++ {
		time.Sleep(time.Millisecond * 1)
		go curlExec(i)
	}
}
