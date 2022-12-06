package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func task(idx int) {
	fmt.Println("now task is", idx)
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i)
	}
	wg.Wait()
	fmt.Println("over")
}
