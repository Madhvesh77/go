package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePageSize struct {
	url string
	size int
}

func main() {
	urls := []string {
		"http://www.apple.com",
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.microsoft.com",
	}
	results := make(chan HomePageSize)

	for _, url := range urls {
		go func (url string) {
			response, error := http.Get(url)
			if error != nil {
				panic(error)
			}
			defer response.Body.Close()

			bs, error := ioutil.ReadAll(response.Body)
			if error != nil {
				panic(error)
			}
			results <- HomePageSize{
				url: url,
				size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

    for range urls {
        result := <-results
        if result.size > biggest.size {
            biggest = result
        }
    }

    fmt.Println("The biggest home page:", biggest.url)
}