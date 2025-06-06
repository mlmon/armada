package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	srcURL := os.Getenv("SRC_URL")
	if srcURL == "" {
		srcURL = "https://httpbin.org/get"
	}

	destPath := os.Getenv("DEST_PATH")
	if destPath == "" {
		destPath = "/tmp/downloaded_file"
	}

	client := &fasthttp.Client{
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(srcURL)
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Set("User-Agent", "armada-exdown/1.0")

	fmt.Printf("Downloading from: %s\n", srcURL)
	fmt.Printf("Saving to: %s\n", destPath)

	err := client.Do(req, resp)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode() != 200 {
		log.Fatalf("HTTP error: %d", resp.StatusCode())
	}

	err = os.WriteFile(destPath, resp.Body(), 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	fmt.Printf("Successfully downloaded %d bytes\n", len(resp.Body()))
	fmt.Printf("Content-Type: %s\n", resp.Header.Peek("Content-Type"))
}