package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {}

// WaitForServer tries to connect to server
// It tries to do for a minute with exponental backoff strategy
// If all trials are failed, return err
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // Success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponental backoff
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
