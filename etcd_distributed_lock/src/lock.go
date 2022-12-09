package lock

import (
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func initEtcdClient() *clientv3.Client {
	var client *clientv3.Client
	var err error

	endpoints := []string{"http://127.0.0.1:2379"}

	// 创建一个 etcd 的客户端
	client, err = clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: 5 * time.Second})
	if err != nil {
		log.Printf("初始化客户端失败: %v\n", err)
		log.Fatal(err)
	}

	return client
}

func Lock(id int, lockName string) {
	client := initEtcdClient()
	defer client.Close()

	// 创建一个 session，如果程序宕机崩溃，etcd可以知道
	s, err := concurrency.NewSession(client)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	// 创建一个 etcd lock
	locker := concurrency.NewLocker(s, lockName)

	log.Printf("id: %v 尝试获取锁: %v", id, lockName)
	locker.Lock()
	log.Printf("id: %v 取得锁: %v", id, lockName)

	// 模拟业务
	time.Sleep(time.Millisecond * 300)

	locker.Unlock()

	log.Printf("id: %v 释放锁: %v", id, lockName)
}
