package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

const (
	EXCHANGE    = "exchange-delay-test"
	QUEUE       = "queue-delay-test"
	ROUTING_KEY = "routingkey-delay-test"
)

func main() {
	fmt.Println("starting rabbitmq consumer....")
	conn, err := amqp.Dial("amqp://test:123456@172.16.10.156:5672/test")

	defer func() {
		if r := recover(); r != nil {
			if v, ok := r.(string); ok {
				fmt.Print(v)
			} else {
				fmt.Printf("error: %v", r)
			}
		}
		fmt.Print("stopping consumer....")
		conn.Close()
	}()

	if err != nil {
		panic("rabbitmq dial failed : " + err.Error())
	}
	fmt.Println("rabbitmq dial succeed")

	ch, err := conn.Channel()
	if err != nil {
		panic("rabbitmq connect failed : " + err.Error())
	}
	fmt.Println("rabbitmq connect succeed")

	chans, err := ch.Consume(QUEUE, "", true, false, true, true, make(map[string]interface{}))

	go func() {
		for msg := range chans {
			fmt.Printf("receive msg : %+v  \n", msg)
		}
	}()

	for {
		fmt.Printf("print 'exit' for exit\n")
		var msg string
		fmt.Scan(&msg)
		fmt.Printf("read from console : %s \n", msg)
		if msg == "exit" {
			break
		}
	}
}
