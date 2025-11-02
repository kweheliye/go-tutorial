package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {

	start := time.Now()
	if resp, err := http.Get(url); err != nil {
		ch <- result{}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {
	stopper := time.After(300 * time.Millisecond)

	result := make(chan result)
	urls := []string{
		"https://www.amazon.com",
		"https://www.google.com",
		"https://www.nytimes.com",
		"https://www.wsj.com",
	}

	start := time.Now()
	for _, url := range urls {
		go get(url, result)
	}

	for range urls {
		select {
		case r := <-result:
			if r.err != nil {
				log.Printf("%s failed: %v", r.url, r.err)
			} else {
				log.Printf("%s took %s", r.url, r.latency)
			}
		case t := <-stopper:
			log.Fatalf("timeout %s", t)
		}
	}

	log.Printf("%d urls took %s", len(urls), time.Since(start))

}
