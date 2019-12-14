/**
 * Author : SimonHuang
 * hgq（huang go MQ）
 * The redis work as Broker;
 */
package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main()  {
	fmt.Println("这个是生产者")

	client := redis.NewClient(&redis.Options{
		Addr:     "10.0.0.103:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	client.Set("SimonHuangTest","OK",0)
	val, err := client.Get("SimonHuangTest").Result();
	if err != nil {
		panic(err)
	}
	fmt.Println("SimonHuangTest", val)

}
