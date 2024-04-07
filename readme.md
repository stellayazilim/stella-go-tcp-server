
# STELLA GO TCP SERVER

This is a simple event driven low level tcp server library in early stage of development and it's not ready for production yet


currently only populates handler after connection finish


## todo
- handle real time data stream
- built in parsers "json, bson, text, binary"
- populate handler on connection establishes and provide events like, data, drain, etc
- concurrent and thread safe
- zero memory allocation on handler registering

## example usage

note: first day implementation it probably changes in future


```go
	s, _ := server.NewServer(":1453")


    /** { "event": "add-token", "data", "dummy data"} **/
	s.Handle("add-token", func(socket *server.Socket) {
		
		fmt.Println("socket connected:", socket.ReadAsString())
        // data received: dummy data
	})

	if err := s.Listen(); err != nil {
		fmt.Println(err)
	}
```
# LICENSE
[AGPL-3.0](./LICENSE)
