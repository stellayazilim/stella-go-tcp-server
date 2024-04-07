package main

import (
	"fmt"
	"github.com/stellayazilim/StellaTCP/pkg/server"
)

func main() {

	s, _ := server.NewServer(":1453")

	s.Handle("add-token", func(socket *server.Socket) {

		fmt.Println("data received:", socket.ReadAsString())

	})

	if err := s.Listen(); err != nil {
		fmt.Println(err)
	}

}
