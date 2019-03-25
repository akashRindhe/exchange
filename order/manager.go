package order

import (
	"fmt"
)

var serverChannel chan Message
var orderBookChannel chan Message

func InitializeOrderManager(comm chan Message) {
	serverChannel = comm
	initializeOrderBooks()
	go fromServer()
	go toServer()
}

func initializeOrderBooks() {

}

func fromServer() {
	for {
		message := <-serverChannel
		fmt.Println("From server: ", message)
	}
}

func toServer() {
	for {
		message := <-orderBookChannel
		serverChannel <- message
	}
}
