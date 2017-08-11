package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handlerChunckedResponse(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "Chunk #%d\n", i)
		flusher.Flush()
		time.Sleep(500 * time.Millisecond)
	}
	flusher.Flush()
}

func main() {
	var httpSrver http.Server
	http.HandleFunc("/chunked", handlerChunckedResponse)
	log.Println("start http listening :18888")
	httpSrver.Addr = ":18888"
	log.Println(httpSrver.ListenAndServe())
}
