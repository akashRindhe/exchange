package main

import (
	"exchange/server"
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number")
		return
	}
	server.Start(arguments[1])
}
