# go-rabbitmq-delay

#### 介绍
rabbitmq延迟队列的golang的Demo; 关于rabbitmq的[延迟队列插件](https://www.rabbitmq.com/community-plugins.html)]

#### 注意
1. 阿里云消息队列RabbitMQ版的mq消费时的consumerTag格式有要求，本项目不能使用的rabbitmq库自动生成
2. 自建的RabbitMQ的延时时长最长为Int的最大值，但是阿里云消息队列RabbitMQ版最大为1天即86400_000毫秒
