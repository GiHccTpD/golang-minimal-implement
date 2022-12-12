package main

import "fmt"

func mapChan(in <-chan interface{}, fn func(interface{}) interface{}) <-chan interface{} {
	out := make(chan interface{})
	if in == nil { // 异常检测
		close(out)
		return out
	}

	go func() { // 启动一个goroutine，实现map的主要逻辑
		defer close(out)
		for v := range in { // 从输入chan读取数据，执行业务逻辑，也就是map操作
			out <- fn(v)
		}
	}()

	return out
}

func reduce(in <-chan interface{}, fn func(r, v interface{}) interface{}) interface{} {
	if in == nil { // 异常检测
		return nil
	}

	out := <-in         // 先读取一个元素
	for v := range in { // 实现reduce的主要逻辑
		out = fn(out, v)
	}

	return out
}

// 生成一个数据流
func asStream(done <-chan struct{}) <-chan interface{} {
	s := make(chan interface{})
	values := []int{1, 2, 3, 4, 5}
	go func() {
		defer close(s)
		for _, v := range values { // 从数组生成
			select {
			case <-done:
				return
			case s <- v:
			}
		}
	}()

	return s
}

func main() {
	in := asStream(nil)

	// map 操作：✖️10
	mapFn := func(v interface{}) interface{} {
		return v.(int) * 10
	}

	// reduce 操作：对map的结果进行累加
	reduceFn := func(r, v interface{}) interface{} {
		return r.(int) + v.(int)
	}

	sum := reduce(mapChan(in, mapFn), reduceFn) // 返回累加结果
	fmt.Println(sum)
}
