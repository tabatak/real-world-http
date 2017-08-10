package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	reader := strings.NewReader("テキスト")
	// file, err := os.Open("main.go")
	// values := url.Values{
	// 	"test": {"value"},
	// }

	// resp, err := http.PostForm("http://localhost:18888", values)
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
