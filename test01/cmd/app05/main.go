package main

import (
	"cache/cache"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	c := cache.New(httpGet)
	flag.Parse()
	urls := flag.Args()
	var n sync.WaitGroup
	for _, url := range urls {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			response, err := c.Get(url)
			if err != nil {
				fmt.Printf("the get error: %v", err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(response.([]byte)))
		}(url)
	}
	n.Wait()
}
func httpGet(url string) (interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Http Get faild. the url %s, err: %v", url, err)
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
