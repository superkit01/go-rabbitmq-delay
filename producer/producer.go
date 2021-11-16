package main

import (
	"fmt"
	"strconv"

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

	defer func() {
		if r := recover(); r != nil {
			if v, ok := r.(string); ok {
				fmt.Print(v)
			}
			fmt.Print(r.(string))
		}
		fmt.Print("stopping producer....")
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

	// EXCHANGE DECLARE
	table := make(map[string]interface{})
	table["x-delayed-type"] = "direct"
	err = ch.ExchangeDeclare(EXCHANGE, "x-delayed-message", true, false, false, true, table)
	if err != nil {
		panic("declare exchange failed : " + err.Error())
	}
	fmt.Println("declare exchange succeed")

	//QUEUE DECLARE
	_, err = ch.QueueDeclare(QUEUE, true, false, false, true, make(map[string]interface{}))
	if err != nil {
		panic("declare queue failed : " + err.Error())
	}
	fmt.Println("declare queue succeed")

	//BINGING DECLARE
	err = ch.QueueBind(QUEUE, ROUTING_KEY, EXCHANGE, true, make(map[string]interface{}))
	if err != nil {
		panic("declare binding failed : " + err.Error())
	}
	fmt.Println("declare binding succeed")

	for i := 0; i <= 10; i++ {
		headers := make(map[string]interface{})
		headers["x-delay"] = 5000
		message := amqp.Publishing{
			Headers: headers,
			Body:    []byte("test" + strconv.Itoa(i)),
		}
		if err = ch.Publish(EXCHANGE, ROUTING_KEY, false, true, message); err != nil {
			fmt.Println("send msg failed ", err.Error())
		} else {
			fmt.Println("send msg succeed ")
		}
	}

}
