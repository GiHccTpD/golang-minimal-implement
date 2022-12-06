# 等待协程退出

```go
➜ wait_goroutine_exit (main) ✗ go run wait_group.go
now task is 1
now task is 9
now task is 6
now task is 7
now task is 8
now task is 2
now task is 0
now task is 3
now task is 5
now task is 4
over
➜ wait_goroutine_exit (main) ✗ go run channel.go
now task is 9
now task is 5
now task is 3
now task is 4
now task is 6
now task is 2
now task is 7
now task is 8
now task is 0
now task is 1
over
➜ wait_goroutine_exit (main) ✗ go run channel2.go
now task is  0
now task is  1
now task is  2
now task is  3
now task is  4
now task is  5
now task is  6
now task is  7
now task is  8
now task is  9
over
➜ wait_goroutine_exit (main) ✗
```

## 方式1:

`sync.WaitGroup`是等待一组协程结束，`sync.WaitGroup`只有3个方法，`Add()`添加一个计数，`Done()`减去一个计数，`Wait()`阻塞直到所有任务完成

## 方式2:

利用缓冲信道`channel`协程之间通讯，其阻塞等待功能实现等待一组协程结束，不能保证其`goroutine`按照顺序执行

## 方式3:

利用无缓冲的信道`channel`协程之间通讯，其阻塞等待功能实现等待一组协程结束，保证了其`goroutine`按照顺序执行