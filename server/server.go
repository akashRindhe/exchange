package server

import (
	"bufio"
	"exchange/order"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

var clientMap = make(map[string]chan order.Message)
var clientMapLock sync.RWMutex

/*Start function initialises the server to listen on a port*/
func Start(port string) {
	count := 1
	PORT := ":" + port
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		log.Fatal(err)
	}

	orderManagerChan := make(chan order.Message, 10000)
	order.InitializeOrderManager(orderManagerChan)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		//create an identifier and a channel and link in map
		clientName := "client" + strconv.Itoa(count)
		count = count + 1
		newClientChannel := make(chan order.Message, 10000)
		clientMapLock.Lock()
		clientMap[clientName] = newClientChannel
		clientMapLock.Unlock()
		go handleConnection(clientName, conn, newClientChannel)
		go toOrderManager(orderManagerChan, newClientChannel)
	}
}

func handleConnection(clientName string, c net.Conn, commChannel chan order.Message) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netdata, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			c.Write([]byte(string("Failed to decode message")))
			continue
		}

		temp := strings.TrimSpace(string(netdata))
		m := make(order.Message)
		m["client"] = clientName
		m["message"] = temp
		commChannel <- m
		fmt.Println("["+clientName+"] sent ", m)
		receivedmessage := <-commChannel
		c.Write([]byte(receivedmessage.ToString()))
	}
	fmt.Println("Ending connection")
	c.Close()
}

func toOrderManager(orderManagerChannel chan order.Message, clientConnectorChannel chan order.Message) {
	for {
		message := <-clientConnectorChannel
		orderManagerChannel <- message
	}
}

func fromOrderManager(orderManagerChannel chan order.Message) {
	for {
		message := <-orderManagerChannel
		clientMapLock.RLock()
		clientMap[message["client"]] <- message
		clientMapLock.RUnlock()
	}
}
