package main

import (
	"net/http"
)

func main() {
	//m := memo.New(getBody)

}
func getBody(url string) (interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, err
}
