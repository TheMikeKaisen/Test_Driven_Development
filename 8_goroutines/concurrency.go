package concurrency

import "time"

type CheckWebsites func(string)bool

type result struct {
	url string
	isPresent bool
}

func  WebsiteChecker(wc CheckWebsites, websites []string) map[string]bool{
	results := make(map[string]bool)
	resultChannel := make(chan result);

	for _, url := range websites{
		go func(){
			time.Sleep(1 * time.Second)
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i:= 0; i<len(websites); i++ {
		r := <- resultChannel
		results[r.url] = r.isPresent
	}

	return results;
}