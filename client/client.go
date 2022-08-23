package main

import (
	"fmt"
	"io/ioutil"
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

func Test3() *http.Response {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8888/hello", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func Request(method string, url string) *http.Response {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res

}

func main() {

	res := Request("GET", "http://localhost:3000/json")
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

}
