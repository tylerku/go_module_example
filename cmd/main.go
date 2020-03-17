package main

import (
	"fmt"
	"io/ioutil"

	"log"
	"net/http"

	"github.com/tylerku/go_module_example/hello"
)

func main() {
	fmt.Println(hello.Hello())

	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
