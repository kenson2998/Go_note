package main

import (
	"fmt"

	"github.com/jack482653/go-ptt-web-craweler/ptt"
)

func main() {
	url := "https://www.ptt.cc/bbs/Beauty/M.1562338754.A.958.html"
	a, err := ptt.NewArticle(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
