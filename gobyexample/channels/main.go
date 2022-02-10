package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Begin")

	//
	resp := SendPostRequest("https://jsonplaceholder.typicode.com/todos/1", []byte{})
	fmt.Println("Response Code", resp.StatusCode)

	orderChan := make(chan *http.Response)
	go SendPostAsync("https://jsonplaceholder.typicode.com/todos/2", []byte{}, orderChan)
	orderResponse := <-orderChan
	defer orderResponse.Body.Close()
	bytes, _ := ioutil.ReadAll(orderResponse.Body)
	fmt.Println(string(bytes))
}

func SendPostAsync(url string, body []byte, rc chan *http.Response) {
	fmt.Println("Sending request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	rc <- response
}
func SendPostRequest(url string, body []byte) *http.Response {
	fmt.Println("Sending request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return response
}
