package main

import (
	"fmt"
)

func thread1(ch chan string, start chan struct{}) {
	var i int
	for range start {
		i++
		ch <- fmt.Sprintf("thread1 - %d", i)
	}
}

func thread2(ch chan string, start chan struct{}) {
	var i int
	for range start {
		i++
		ch <- fmt.Sprintf("thread2 - %d", i)
	}
}

func main() {
	var (
		th1   = make(chan string)
		th2   = make(chan string)
		start = make(chan struct{})
	)

	go thread1(th1, start)
	go thread2(th2, start)

	for {
		var input string
		fmt.Scanln(&input)
		if input == "t" {
			start <- struct{}{}
			start <- struct{}{}
			fmt.Println(<-th1)
			fmt.Println(<-th2)
		}
	}
}
