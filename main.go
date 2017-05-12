package main

import (
	"btspider/Init"
	. "btspider/Global"
	"fmt"
)

func init() {
	Init.Init("develpment")
}

func main() {
	var request = Request{
		Url: "www.baidu.com",
	}

	result := Spider.Open(request).Resault()
	fmt.Println(string(result.Body))
}
