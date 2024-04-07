package server

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
	"net"
)

type SocketHandlerFunc func(socket *Socket)

type Server interface {
	Handle(event string, handler SocketHandlerFunc)
	Listen() error
}

type incomingEvent struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

type ISocket interface {
	GetConnection() net.Conn
	ReadAsString() (string, error)
}
type Socket struct {
	Id         uuid.UUID
	connection net.Conn
	*incomingEvent
	data  []byte
	event string
}

type server struct {
	listener net.Listener
	handlers map[string]SocketHandlerFunc
}

func NewServer(addr string) (Server, error) {
	handlers := make(map[string]SocketHandlerFunc)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	server := &server{
		listener: listener,
		handlers: handlers,
	}

	return server, nil
}

func (s *server) Handle(event string, handler SocketHandlerFunc) {

	s.handlers[event] = handler
}

func (s *server) Listen() error {

	for {
		con, err := s.listener.Accept()

		if err != nil {
			return err
		}
		go parseIncomingEvents(s, con)

		if err != nil {
			return err
		}

	}

}

// @todo parse realtime

func parseIncomingEvents(s *server, conn net.Conn) {

	b := new(bytes.Buffer)
	ie := new(incomingEvent)

	for {
		_, err := io.CopyN(b, conn, 16)

		if err != nil {
			if err == io.EOF {
				err := conn.Close()
				if err != nil {
					log.Printf("Error closing connection: %v", err)
					break
				}
			} else {
				log.Printf("Error closing connection: %v", err)
				break
			}
		}

		err = json.Unmarshal(b.Bytes(), &ie)

		if err != nil {
			continue
		}

		sock := Socket{
			Id:         uuid.New(),
			data:       []byte(ie.Data),
			connection: conn,
			event:      ie.Event,
		}

		if handler, ok := s.handlers[string(sock.event)]; ok {
			handler(&sock)
		}

	}

}

func (i *incomingEvent) GetEvent() string {
	return string(i.Event)
}

func (i *incomingEvent) SetRaw(b []byte) {
	panic("not implemented")
}

func (s *Socket) GetConnection() net.Conn {
	return s.connection
}

func (s *Socket) ReadAsString() string {

	return string(s.data)
}

func (s *Socket) ReadAsBytes() []byte {
	panic("not implemented")
}

func (s *Socket) ReadAsJosn(dest interface{}) error {
	panic("not implemented")
}

func (s *Socket) ReadAsBson(dest interface{}) error {
	panic("not implemented")
}
