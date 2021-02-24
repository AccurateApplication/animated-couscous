package waitsorserver

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) error { // reports error if attempts fail, if succed, returns error=nil
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server no response: %s , retrying..", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server: %s failed to response after: %s", url, timeout)
}
