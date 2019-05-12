package main
// got it from https://pocketgophers.com/limit-concurrent-use/
import (
	"log"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Ltime)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 3) // we have buffer size 3. And all other will wait.
	for i := 0; i < 9; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			semaphore <- struct{}{} // Lock
			defer func() {
				<-semaphore // Unlock
			}()

			time.Sleep(time.Second)
			log.Println(i)
		}(i)
	}
	wg.Wait()
}
