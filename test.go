package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

const (
	EXCHANGE    = "exchange-delay-test"
	QUEUE       = "queue-delay-test"
	ROUTING_KEY = "routingkey-delay-test"
)

func main() {
	conn, err := amqp.Dial("amqp://test:123456@172.16.0.156:5672/test")

	defer func() {
		if r := recover(); r != nil {
			if v, ok := r.(string); ok {
				fmt.Print(v)
			}
			fmt.Print(r.(string))
		}
		fmt.Print("stopping test....")
		conn.Close()
	}()

	if err != nil {
		panic("rabbitmq dial failed " + err.Error())
	}
	fmt.Println("rabbitmq dial succeed")

	ch, err := conn.Channel()
	if err != nil {
		panic("rabbitmq connect failed : " + err.Error())
	}
	fmt.Println("rabbitmq connect succeed")

	for i := 0; i <= 10; i++ {
		table := make(map[string]interface{})
		table["x-delay"] = 5000
		message := amqp.Publishing{
			Headers: table,
			Body:    []byte("test" + strconv.Itoa(i)),
		}
		if err = ch.Publish(EXCHANGE, ROUTING_KEY, false, false, message); err != nil {
			fmt.Println("send msg failed ", err.Error())
		} else {
			fmt.Println("send msg succeed ")
		}

	}

	time.Sleep(time.Duration(10) * time.Second)

}
