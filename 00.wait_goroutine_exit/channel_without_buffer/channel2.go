package main

import "fmt"

var ch2 = make(chan int)

func taskCh2(idx int) {
	fmt.Println("now task is ", idx)
	<-ch2
}

func main() {
	for i := 0; i < 10; i++ {
		go taskCh2(i)
		ch2 <- i
	}
	fmt.Println("over")
}
