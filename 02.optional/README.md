# 选项模式初始化参数

示例来自[ants](https://github.com/panjf2000/ants)
```bash
➜ optional (main) ✗ go run main.go
o1: &{ExpiryDuration:0s PreAlloc:false MaxBlockingTasks:0 Nonblocking:false PanicHandler:<nil> Logger:<nil>}
o2: &{ExpiryDuration:0s PreAlloc:true MaxBlockingTasks:0 Nonblocking:false PanicHandler:<nil> Logger:<nil>}
o3: &{ExpiryDuration:5s PreAlloc:true MaxBlockingTasks:0 Nonblocking:false PanicHandler:<nil> Logger:<nil>}
```
