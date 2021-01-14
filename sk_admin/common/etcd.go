package common

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

const (
	etcdUrl = "127.0.0.1:2379"
)

func connectEtcd() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdUrl},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("connect to etcd failed, err:%v\n", err)
		return nil, err
	}

	return cli, nil
}

func EtcdPut(seckillIdStr string, seckillJson string) error {
	cli, err := connectEtcd()
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, seckillIdStr, seckillJson)
	cancel()
	if err != nil {
		log.Printf("put to etcd failed, err:%v\n", err)
		return err
	}

	return nil
}

func EtcdGet(seckillIdStr string) (string, error) {
	cli, err := connectEtcd()
	if err != nil {
		return "", err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, seckillIdStr)
	cancel()
	if err != nil {
		log.Printf("get from etcd failed, err:%v\n", err)
		return "", err
	}

	if resp.Kvs == nil {
		log.Printf("etct seckillId [%s] value is empty ", seckillIdStr)
		return "", nil
	}

	value := string(resp.Kvs[0].Value)
	log.Printf("etct seckillId [%s] value: [%s]", seckillIdStr, value)

	return value, nil
}

func EtcdDelete(seckillIdStr string) error {
	cli, err := connectEtcd()
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Delete(ctx, seckillIdStr)
	cancel()
	if err != nil {
		log.Printf("delete seckillIdStr:[%s] from etcd failed, err:%v\n", seckillIdStr, err)
		return err
	}

	return nil
}
