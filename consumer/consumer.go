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
	fmt.Println("starting rabbitmq producer....")
	conn, err := amqp.Dial("amqp://test:123456@172.16.0.156:5672/test")
	if err != nil {
		fmt.Println("rabbitmq dial failed : ", err.Error())
	} else {
		fmt.Println("rabbitmq dial succeed")
	}

	defer func() {
		fmt.Print("stopping producer....")
		conn.Close()
	}()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("rabbitmq connect failed : ", err.Error())
	} else {
		fmt.Println("rabbitmq connect succeed")
	}

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
