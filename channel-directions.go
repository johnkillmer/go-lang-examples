package main

import "fmt"

func ping(pings chan<- string, msg string) {
	fmt.Println("ping:" + msg)
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- "reply:" + msg
}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	for i := 0; i < 5; i++ {
		ping(pings, "passed message")
		pong(pings, pongs)
		fmt.Println(<-pongs)
	}

}
