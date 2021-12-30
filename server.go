package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "안녕하세요! %s에 오신 것을 환경합니다!", r.URL.Path[1:])
}
