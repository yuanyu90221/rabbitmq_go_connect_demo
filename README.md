# rabbitmq with golang

## introduction

This repository is about use RabbitMQ to implement a WorkQueue with addTask


## RabbitMQ with golang
## Consumer
This consumer create a go channel to consume the message from rabbitmq with queuename = 'add'

And use a goroutine to receive data from consume channel

And sum the data from channel

## Producer
This producer publish a the data to the rabbitmq with queuename = 'add'

the data contains 2 integer numbers, Number1, Number2

## Advantage of this pub/sub pattern
1 All the request will be queue on ampq until the consumer consume, less data loss
2 consumer and producer don't have to be connected directly and can be asynchronize handle
