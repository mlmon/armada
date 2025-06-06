package main

import (
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	client := &fasthttp.Client{
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	urls := []string{
		"https://httpbin.org/get",
		"https://httpbin.org/status/200",
		"https://httpbin.org/json",
	}

	for _, url := range urls {
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)

		req.SetRequestURI(url)
		req.Header.SetMethod(fasthttp.MethodGet)
		req.Header.Set("User-Agent", "armada-exdown/1.0")

		fmt.Printf("Making request to: %s\n", url)

		err := client.Do(req, resp)
		if err != nil {
			log.Printf("Error making request to %s: %v", url, err)
			continue
		}

		fmt.Printf("Status: %d\n", resp.StatusCode())
		fmt.Printf("Content-Type: %s\n", resp.Header.Peek("Content-Type"))
		fmt.Printf("Body length: %d bytes\n", len(resp.Body()))
		fmt.Println("---")
	}

	fmt.Println("FastHTTP client example completed")
}