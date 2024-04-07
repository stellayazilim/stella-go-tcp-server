package helpers

import (
	"bufio"
	"fmt"
	"net"
)

func ParseIncomingConnectionSchema(conn net.Conn, ch chan error, ct chan<- string) {

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			err := conn.Close()

			ch <- err
			return
		}
		fmt.Printf("Message incoming: %s", string(message))
		_, err = conn.Write([]byte("Message received.\n"))

		ct <- message
		if err != nil {
			ch <- err
			return
		}
	}
}
