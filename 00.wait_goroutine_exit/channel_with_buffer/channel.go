package main

import "fmt"

var ch = make(chan int, 10)

func taskCh(idx int) {
	fmt.Println("now task is", idx)
	ch <- idx
}

func main() {
	for i := 0; i < 10; i++ {
		go taskCh(i)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	fmt.Println("over")
}
