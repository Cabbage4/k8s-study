package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rsp http.ResponseWriter, req *http.Request) {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			rsp.Write([]byte(err.Error()))
			return
		}

		log.Println(fmt.Sprintf("hello world:\n%s", b))

		rsp.Write([]byte(fmt.Sprintf("hello world:\n%s", b)))
	})

	http.ListenAndServe(":8082", nil)
}
