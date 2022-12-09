package main

import (
	"sync"

	lock "etcd_distributed_lock/src"
)

func main() {
	lockName := "locker-test"
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(item int) {
			defer wg.Done()
			lock.Lock(item, lockName)
		}(i)
	}
	wg.Wait()
}
