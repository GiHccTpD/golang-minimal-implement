# ETCD分布式锁

```bash
➜ etcd_distributed_lock (main) ✗ go run main.go
2022/12/10 00:11:55 id: 2 尝试获取锁: locker-test
2022/12/10 00:11:55 id: 9 尝试获取锁: locker-test
2022/12/10 00:11:55 id: 7 尝试获取锁: locker-test
2022/12/10 00:11:55 id: 5 尝试获取锁: locker-test
2022/12/10 00:11:55 id: 8 尝试获取锁: locker-test
2022/12/10 00:11:55 id: 4 尝试获取锁: locker-test
2022/12/10 00:11:55 id: 6 尝试获取锁: locker-test
2022/12/10 00:11:55 id: 3 尝试获取锁: locker-test
2022/12/10 00:11:55 id: 0 尝试获取锁: locker-test
2022/12/10 00:11:55 id: 1 尝试获取锁: locker-test
2022/12/10 00:11:55 id: 9 取得锁: locker-test
2022/12/10 00:11:56 id: 9 释放锁: locker-test
2022/12/10 00:11:56 id: 4 取得锁: locker-test
2022/12/10 00:11:56 id: 4 释放锁: locker-test
2022/12/10 00:11:56 id: 8 取得锁: locker-test
2022/12/10 00:11:56 id: 8 释放锁: locker-test
2022/12/10 00:11:56 id: 1 取得锁: locker-test
2022/12/10 00:11:57 id: 1 释放锁: locker-test
2022/12/10 00:11:57 id: 0 取得锁: locker-test
2022/12/10 00:11:57 id: 0 释放锁: locker-test
2022/12/10 00:11:57 id: 5 取得锁: locker-test
2022/12/10 00:11:57 id: 5 释放锁: locker-test
2022/12/10 00:11:57 id: 6 取得锁: locker-test
2022/12/10 00:11:57 id: 6 释放锁: locker-test
2022/12/10 00:11:57 id: 2 取得锁: locker-test
2022/12/10 00:11:58 id: 2 释放锁: locker-test
2022/12/10 00:11:58 id: 3 取得锁: locker-test
2022/12/10 00:11:58 id: 3 释放锁: locker-test
2022/12/10 00:11:58 id: 7 取得锁: locker-test
2022/12/10 00:11:58 id: 7 释放锁: locker-test
```
