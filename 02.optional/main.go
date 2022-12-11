package main

import (
	"demo/lib"
	"fmt"
	"time"
)

func main() {
	o := lib.New()
	fmt.Printf("o1: %+v\n", o)

	o = lib.New(lib.WithPreAlloc(true))
	fmt.Printf("o2: %+v\n", o)

	o = lib.New(lib.WithPreAlloc(true), lib.WithExpiryDuration(5*time.Second))
	fmt.Printf("o3: %+v\n", o)
}
