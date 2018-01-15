package main

import (
	"fmt"
	"sync"
)

func main() {
	numChannel1 := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			numChannel1 <- fmt.Sprint("numChannel1:", i)
		}
		close(numChannel1)
	}()

	for number := range numChannel1 {
		fmt.Println(number)
	}

	/* using waitgroup */

	numChannel2 := make(chan string)
	var waitGroup2 sync.WaitGroup
	waitGroup2.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			numChannel2 <- fmt.Sprint("numChannel2:1:", i)
		}
		waitGroup2.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			numChannel2 <- fmt.Sprint("numChannel2:2:", i)
		}
		waitGroup2.Done()
	}()

	go func() {
		waitGroup2.Wait()
		close(numChannel2)
	}()

	for number := range numChannel2 {
		fmt.Println(number)
	}

	/* using semaphores */

	numChannel3 := make(chan string)
	doneChannel3 := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			numChannel3 <- fmt.Sprint("numChannel3:1:", i)
		}
		doneChannel3 <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			numChannel3 <- fmt.Sprint("numChannel3:2:", i)
		}
		doneChannel3 <- true
	}()

	go func() {
		<-doneChannel3
		<-doneChannel3
		close(numChannel3)
	}()

	for number := range numChannel3 {
		fmt.Println(number)
	}

	/* multiple go routines polling single channel */

	numChannel4 := make(chan string)
	doneChannel4 := make(chan bool)

	go func() {
		for i := 0; i < 1000; i++ {
			numChannel4 <- fmt.Sprint("numChannel4:", i)
		}
		close(numChannel4)
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for number := range numChannel4 {
				fmt.Println(number)
			}
			doneChannel4 <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-doneChannel4
	}

	/* passing and returning channels */

	incrementorChannel := incrementor()
	for number := range sum(incrementorChannel) {
		fmt.Println(number)
	}

}

func incrementor() chan int {
	channel := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		close(channel)
	}()
	return channel
}

func sum(inputChannel chan int) chan int {
	channel := make(chan int)
	go func() {
		var sum int
		for number := range inputChannel {
			sum += number
		}
		channel <- sum
		close(channel)
	}()
	return channel
}
