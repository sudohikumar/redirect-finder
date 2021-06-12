package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var httpClient *http.Client

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No URLs supplied")
	}

	c := make(chan []string)
	for _, v := range os.Args[1:] {
		go makeRequest(http.MethodGet, v, c)
	}

	fmt.Println()
	for range os.Args[1:] {
		v := <-c
		fmt.Printf("URL: %v - Redirected URL: %v\n", v[0], v[1])
	}
}

func makeRequest(method, url string, c chan []string) {
	req, err := getRequestClient(method, url)
	if err != nil {
		c <- []string{url, err.Error()}
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := getHTTPClient().Do(req)
	if err != nil {
		c <- []string{url, err.Error()}
		return
	}

	if resp.Request.URL != nil {
		c <- []string{url, resp.Request.URL.Host}
	} else {
		c <- []string{url, ""}
	}
}

func getHTTPClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return httpClient
}

func getRequestClient(method, url string) (*http.Request, error) {
	return http.NewRequest(method, url, nil)
}
