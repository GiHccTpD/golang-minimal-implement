//go:build linux || darwin
// +build linux darwin

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var signalFuncList []func()

func init() {
	signalFuncList = make([]func(), 0, 10)
}

func AddSignalFunc(signalFunc func()) {
	signalFuncList = append(signalFuncList, signalFunc)
}

func WaitSignal() {
	fmt.Println("wait signal")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP, syscall.SIGKILL)
	select {
	case a := <-c:
		fmt.Println("接受到退出信号: ", a.String())
		for _, s := range signalFuncList {
			s()
		}
		return
	}
}

func main() {
	fmt.Println("main run~")

	AddSignalFunc(func() {
		fmt.Println("exit 1")
	})
	AddSignalFunc(func() {
		fmt.Println("exit 2")
	})
	WaitSignal()
}
