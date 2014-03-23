package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"time"
)

func main() {

	const (
		HeartBeatTime = 30
	)

	nc, err := net.Dial("tcp", ":6969")
	if err != nil {
		fmt.Println(err)
	}
	defer nc.Close()
	// nc.SetReadDeadline(time.Now().Add(4 * time.Second))

	c := redis.NewConn(nc, 0, 0)

	//订阅频道
	c.Send("sub", "Channel", HeartBeatTime)
	c.Flush()

	// 定时请求心跳
	ticker := time.NewTicker(time.Second * HeartBeatTime)
	go func() {
		for t := range ticker.C {
			nc.Write([]byte("h"))
			fmt.Println("Heartbeat send", t)
		}
	}()

	for {

		reply, err := c.Receive()
		if err != nil {
			fmt.Println(err)
		}

		if reply == "h" {
			fmt.Println("Heartbeat received")
		} else if reply != nil {
			reply_str, err := redis.String(reply, err)
			fmt.Printf("%#v\n", reply_str)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}
