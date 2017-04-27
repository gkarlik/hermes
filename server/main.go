package main

import (
	"fmt"

	"github.com/gkarlik/hermes"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	socket, _ := zmq.NewSocket(zmq.ROUTER)
	defer socket.Close()
	socket.Bind("tcp://*:8080")

	for {
		addr, err := socket.Recv(0)
		if err != nil {
			panic(err)
		}

		msg, err := socket.RecvMessage(0)
		if err != nil {
			panic(err)
		}

		m, err := hermes.ParseMessage(msg)
		if err == nil {
			fmt.Printf("message: %v", m)

			fmt.Printf("sending message: %s\r\n", m.Body)
			if _, err := socket.SendMessage(addr, m.ToStringArray()); err != nil {
				panic(err)
			}
		}
	}
}
