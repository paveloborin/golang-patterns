package primitives

import (
	"sync"
	"fmt"
)

func mutex() {

	var count int
	var lock sync.Mutex

	increment :=func(){
		//Lock is the request on exclusive use of critical section
		lock.Lock()
		//Always call Unlock within a defer statement. This is a very common idiom.
		defer lock.Unlock()

		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func(){
		lock.Lock()

		defer lock.Unlock()

		count--
		fmt.Printf("Decrementing: %d\n", count)
	}


	//Increment
	var arithmetic sync.WaitGroup
	for i:=0; i<5; i++ {
		arithmetic.Add(1)
		go func(){
			defer arithmetic.Done()
			increment()
		}()
	}

	//Decrement
	for i:=0; i<=5; i++ {
		arithmetic.Add(1)
		go func(){
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete")
}
