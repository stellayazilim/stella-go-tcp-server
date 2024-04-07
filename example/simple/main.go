package main

import (
	"fmt"
	"github.com/stellayazilim/StellaTCP/internal/server"
)

func main() {

	s, _ := server.NewServer(":1453")

	s.Handle("add-token", func(socket *server.Socket) {

		fmt.Println("socket connected:", socket.ReadAsString())

	})

	if err := s.Listen(); err != nil {
		fmt.Println(err)
	}

}
