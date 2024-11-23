package main

import ( "fmt"
"time"
)

func main() {
	var channel1 chan string = make(chan string)

	go pinger(channel1)
	go ponger(channel1)
	go printer(channel1)

	var input string
	fmt.Scanln(&input)
}

func pinger(c chan<- string) {
	for i:=0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string ) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println((msg))
		time.Sleep(time.Second * 1)
	}
}