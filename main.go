package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	numOfWorkers := flag.Int("w", 4, "number of workers")
	rangeOfInts := flag.Int("r", 10000, "upper limit of range of ints")
	flag.Parse()

	process(*numOfWorkers, intGenerator(*rangeOfInts))

	fmt.Println("all goroutines processed")
}

// intGenerator generates random int [1, n]
func intGenerator(n int) <-chan int {
	intStream := make(chan int, n)
	go func() {
		rand.Seed(time.Now().UnixNano())
		p := rand.Perm(n + 1)
		for _, r := range p {
			// eleminates 0
			if r == 0 {
				continue
			}
			intStream <- r
		}
		close(intStream)
	}()
	return intStream
}

// process represents a worker pool
func process(numOfWorkers int, in <-chan int) {
	var wg sync.WaitGroup
	wg.Add(numOfWorkers)
	for i := 0; i < numOfWorkers; i++ {
		go doJob(in, &wg)
	}
	wg.Wait()
}

func doJob(intStream <-chan int, wg *sync.WaitGroup) {
	for n := range intStream {
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
	wg.Done()
}
