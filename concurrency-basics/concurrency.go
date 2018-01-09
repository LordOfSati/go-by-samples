package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var waitGroup sync.WaitGroup
var mutex sync.Mutex
var counter1 int
var counter2 int
var counter3 int64

func main() {
	waitGroup.Add(8)
	go printMethodA()
	go printMethodB()
	/* incrementor1 with race condition - check by executing go run -race concurrency.go */
	go incrementor1()
	go incrementor1()
	/* avoiding race condition using mutex - check by executing go run -race concurrency.go */
	go incrementor2()
	go incrementor2()
	/* avoiding race condition using atomic - check by executing go run -race concurrency.go */
	go incrementor3()
	go incrementor3()
	waitGroup.Wait()
	fmt.Println("Value of counter1", counter1) /* not always 50, due to race condition */
	fmt.Println("Value of counter2", counter2)
	fmt.Println("Value of counter3", counter3)
}

func printMethodA() {
	for i := 0; i < 10; i++ {
		fmt.Println("Method A", i)
		time.Sleep(time.Millisecond * 5)
	}
	waitGroup.Done()
}

func printMethodB() {
	for i := 0; i < 10; i++ {
		fmt.Println("Method B", i)
		time.Sleep(time.Millisecond * 10)
	}
	waitGroup.Done()
}

func incrementor1() {
	for i := 0; i < 25; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter1++
	}
	waitGroup.Done()
}

func incrementor2() {
	for i := 0; i < 25; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
		counter2++
		mutex.Unlock()
	}
	waitGroup.Done()
}

func incrementor3() {
	for i := 0; i < 25; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter3, 1)
	}
	waitGroup.Done()
}
