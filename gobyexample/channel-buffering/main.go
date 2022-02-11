package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Begin")

	orderChan := make(chan *http.Response, 2)
	go SendPostAsync("https://jsonplaceholder.typicode.com/todos/1", []byte{}, orderChan)
	orderResponse := <-orderChan
	defer orderResponse.Body.Close()
	bytes, _ := ioutil.ReadAll(orderResponse.Body)
	fmt.Println(string(bytes))

	go SendPostAsync("https://jsonplaceholder.typicode.com/todos/2", []byte{}, orderChan)
	orderResponse2 := <-orderChan
	defer orderResponse2.Body.Close()
	bytes2, _ := ioutil.ReadAll(orderResponse2.Body)
	fmt.Println(string(bytes2))

}

func SendPostAsync(url string, body []byte, rc chan *http.Response) {
	fmt.Println("Sending async request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	rc <- response
}
