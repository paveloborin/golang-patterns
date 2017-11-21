package primitives

import (
	"sync"
	"time"
	"fmt"
)

func cond() {

	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		//This is one of two methods that the Cond type provides for notifying goroutines blocked on a Wait call that the condition has been triggered
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct {
		}{})

		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
