/**
 * Author : SimonHuang
 * hgq（huang go MQ）
 * The redis work as Broker;
 */
package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

func main() {
	fmt.Println("这个是生产者")

	client := redis.NewClient(&redis.Options{
		Addr:     "10.0.0.103:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	pubsub := client.PSubscribe("mychannel*")
	fmt.Println(pubsub)

	for i := 0; i < 100; i++ {
		client.Publish("mychannel1", i)
	}

	client.Close()
}
