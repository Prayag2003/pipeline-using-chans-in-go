package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// NOTE: Control the amount of data we take from the stream
func takeData[T any, P any](done <-chan P, stream <-chan T, num int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			// reading the value from the 'stream' and writing it to the 'taken' stream
			case taken <- <-stream:
			}
		}
	}()
	return taken
}

// NOTE: Slow pipeline
func isPrimeStream(done <-chan int, randStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		for i := 2; i < randomInt; i++ {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomInt := <-randStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()
	return primes
}

func fanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup

	fannedInStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:

			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream

}

func main() {

	start := time.Now()
	done := make(chan int)
	defer close(done)

	randNumber := func() int { return rand.Intn(1e9) }
	randStream := repeatFunc(done, randNumber)
	// primes := isPrimeStream(done, randStream)

	// for random := range takeData(done, primes, 10) {
	// 	fmt.Println(random)
	// }

	numOfCPU := runtime.NumCPU()
	primeChannels := make([]<-chan int, numOfCPU)

	// Fan Out
	for i := 0; i < numOfCPU; i++ {
		primeChannels[i] = isPrimeStream(done, randStream)
	}

	// Fan In
	fannedInStream := fanIn(done, primeChannels...)

	for random := range takeData(done, fannedInStream, 10) {
		fmt.Println(random)
	}

	fmt.Println(time.Since(start))
}

// Repeatedly calls a function that is passed into it
// Return a read only channel
func repeatFunc[T any, P any](done <-chan P, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)

		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}
