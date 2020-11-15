package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	simpleRequest()
}

func simpleRequest() {
	url := "http://ergast.com/api/f1/drivers.json"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
