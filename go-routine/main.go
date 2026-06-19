package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello (message string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("Hello", message)
}

func main (){
	var wg sync.WaitGroup

	totalJobs := 5;

	for i := 1; i <= totalJobs; i++ {
		wg.Add(1)
		go sayHello(fmt.Sprintf("JOB %d", i), time.Second, &wg)
	}

	wg.Wait()
	fmt.Println("All jobs completed")
}

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func sayHello (message string, delay time.Duration, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(delay)
// 	fmt.Println("Hello", message)
// }

// func main (){
// 	var wg sync.WaitGroup

// 	// 1. Add outside of the goroutine
// 	// 2. You must dcrease the counter in the goroutine using defer wg.Done() to ensure that it is called even if the goroutine panics or returns early.
// 	// 3. Do not forget to call wg.Wait() in the main goroutine to block until all goroutines have completed their work.
// 	// 4. Allways pass a reference/pointer to the WaitGroup to the goroutines so that they can access and modify the counter correctly.

// 	wg.Add(4)

// 	fmt.Println("Hello from main Goroutine starting")
// 	go sayHello("Goroutine 1", time.Second, &wg)
// 	go sayHello("hellow world wait", time.Second, &wg)
// 	go sayHello("Goroutine 2", 2*time.Second, &wg)
// 	go sayHello("Gouroutine3", 5*time.Second, &wg)
// 	wg.Wait()
// 	fmt.Println("Hello from main Goroutine Last")
// }
