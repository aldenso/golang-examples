package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func worker(url string, pool chan struct{}, wg *sync.WaitGroup, start time.Time) {
	startUrl := time.Now()
	defer func() {
		<-pool
	}()
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}

	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Request timed out for", url, "wasted", time.Since(startUrl))
		} else {
			log.Println(err)
		}
		return
	}
	defer resp.Body.Close()

	fmt.Println("url:", url, "code:", resp.StatusCode, "took:", time.Since(startUrl), "took in total:", time.Since(start))

}

func main() {
	start := time.Now()
	urls := []string{
		"https://example.com",
		"https://www.ibm.com",
		"https://google.com",
		"https://github.com",
		"https://yahoo.com",
		"https://www.oracle.com",
		"https://www.youtube.com",
		"https://bing.com",
	}

	maxWorkers := 3
	var wg sync.WaitGroup
	pool := make(chan struct{}, maxWorkers)

	for _, url := range urls {
		pool <- struct{}{}
		wg.Add(1)
		go worker(url, pool, &wg, start)
	}

	wg.Wait()
	fmt.Println("Done in:", time.Since(start))
}
