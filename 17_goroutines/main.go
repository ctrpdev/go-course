package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan) // solo tiene sentido cuando sabes que funcion demora mas
}

func main() {
	done := make(chan bool)

	// dones := make([]chan bool, 4)
	// for i := range dones {
	// 	dones[i] = make(chan bool)
	// }
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How... are... you?...", done)
	go greet("I hope you're liking the course!", done)

	// <-done
	// <-done
	// <-done
	// <-done

	// for _, done := range dones {
	// 	<-done
	// }

	for range done {
		// fmt.Println(doneChan)
	}

}
