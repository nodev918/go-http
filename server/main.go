package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
	fmt.Fprintf(w, "hello")
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Println("homee")
	fmt.Fprintf(w, "homeeee")
}

func json(w http.ResponseWriter, req *http.Request) {
	fmt.Println("json")
	data := []byte(`{"k":"v"}`)
	fmt.Fprintf(w, string(data))
}

func main() {
	//s := http.Server{}
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/json", json)

	http.ListenAndServe(":8888", nil)
}
