package main

import (
	"fmt"

	"../config"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("starting rabbitmq consumer....")
	conn, err := amqp.Dial("amqp://" + config.USERNAME + ":" + config.PASSWORD + "@" + config.ADDRESS + "/" + config.VHOST)

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

	// chans, err := ch.Consume(QUEUE, "", true, false, true, true, make(map[string]interface{}))

	// go func() {
	// 	for msg := range chans {
	// 		fmt.Printf("receive msg : %+v  \n", msg)
	// 	}
	// }()

	msgs, err := ch.Consume(
		config.QUEUE, // queue
		"",
		// "ctag-C:\\Users\\YDY00076\\AppData\\Local\\Temp\\___go_build_go_rabbitmq_consumer.exe-1", // consumer
		// "test",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s \n", d.Body)
		}
	}()

	//  // CONSUME BY GET
	// for {
	// 	message, ok, err := ch.Get(QUEUE, true)
	// 	if err != nil {
	// 		fmt.Printf("%+v", err.Error())
	// 		time.Sleep(time.Duration(2) * time.Second)
	// 	} else {
	// 		fmt.Printf("message:%+v", message)
	// 		fmt.Printf("ok: %+v", ok)
	// 	}
	// }

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
