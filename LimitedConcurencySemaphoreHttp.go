package main

// got it from https://pocketgophers.com/limit-concurrent-use/
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	log.SetFlags(log.Ltime)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 13) // we have buffer size 3. And all other will wait.
	out := make(map[int]string)
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			semaphore <- struct{}{} // Lock
			defer func() {
				<-semaphore // Unlock
			}()

			resp, err := http.Get("https://google.com")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			out[i] = string(body)
			log.Println(i)
		}(i)
	}
	wg.Wait()

}
