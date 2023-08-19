package chat

type Hub struct {
	Rooms      map[int]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[int]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

// func (hub *Hub) Run() {
// 	for {
// 		select {
// 		case cl := <-hub.Register:

// 		case cl := <-hub.Unregister:

// 		case cl := <-hub.Broadcast:

// 		}
// 	}
// }
