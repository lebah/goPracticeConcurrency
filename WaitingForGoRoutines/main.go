// Write your answer here, and then test your code.
// Your job is to convert MultiURLTime to run concurrently.
package DraftMain

import (
	"io"
	"log"
	"net/http"
	"time"
)

// MutliURLTimes calls URLTime for every URL in URLs.
func MultiURLTime(urls []string) {
	for _, url := range urls {
		URLTime(url)
	}
}

// URLTime checks how much time it takes url to respond.
func URLTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error: %q - %s", url, err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("error: %q - bad status - %s", url, resp.Status)
		return
	}
	// Read body
	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		log.Printf("error: %q - %s", url, err)
		return
	}

	duration := time.Since(start)
	log.Printf("info: %q - %v", url, duration)
}

func main() {
	start := time.Now()

	urls := []string{
		"http://localhost:8080/200",
		"http://localhost:8080/100",
		"http://localhost:8080/50",
	}

	MultiURLTime(urls)

	duration := time.Since(start)
	log.Printf("%d URLs in %v", len(urls), duration)
}
