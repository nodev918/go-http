package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Test1() *http.Response {
	req, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	// fmt.Println(res)
	// fmt.Println("")
	// fmt.Println(reflect.TypeOf(res))
	// fmt.Println(res.StatusCode)

	return res
}

func Test2() {
	res, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("%v: ", res.StatusCode)
	}

}

func main() {
	res := Test1()
	fmt.Println(res.StatusCode)
	// Test2()
}
