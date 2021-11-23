package main

import (
	"fmt"
	"go-rabbitmq/config"
	"strconv"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("starting rabbitmq producer....")
	conn, err := amqp.Dial("amqp://" + config.USERNAME + ":" + config.PASSWORD + "@" + config.ADDRESS + "/" + config.VHOST)

	defer func() {
		if r := recover(); r != nil {
			if v, ok := r.(string); ok {
				fmt.Print(v)
			}
			// fmt.Print(r.(string))
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

	if err = ch.ExchangeDelete(config.EXCHANGE, false, false); err != nil {
		panic("exchange delete failed" + err.Error())
	}
	if count, err := ch.QueueDelete(config.QUEUE, false, false, false); err != nil {
		panic("queue delete failed" + err.Error())
	} else {
		fmt.Printf("deleted %s queues   \n", strconv.Itoa(count))
	}

}
