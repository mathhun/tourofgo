package main

import (
	"fmt"
	"time"
)

//func main() {
//	c := make(chan int, 2)
//	c <- 1
//	c <- 2
//	c <- 3
//	fmt.Println(<-c)
//	fmt.Println(<-c)
//	fmt.Println(<-c)
//}

var c = make(chan int, 5)

func main() {
	go worker(1)
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println(i)
	}
}

func worker(id int) {
	for {
		_ = <-c
		time.Sleep(time.Second)
	}
}
