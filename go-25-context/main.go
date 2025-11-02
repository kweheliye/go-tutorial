package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {

	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {
	result := make(chan result)
	urls := []string{
		"https://www.amazon.com",
		"https://www.google.com",
		"https://www.nytimes.com",
		"https://www.wsj.com",
		"https://10.255.255.1", // an unroutable IP → guaranteed timeout
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, url := range urls {
		go get(ctx, url, result)
	}

	for range urls {
		select {
		case r := <-result:
			if r.err != nil {
				log.Printf("%s failed: %v", r.url, r.err)
			} else {
				log.Printf("%s took %s", r.url, r.latency)
			}

		}
	}

	log.Print("Done")

}
