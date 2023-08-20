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

func (hub *Hub) Run() {
	for {
		select {

		case cl := <-hub.Register:
			if _, ok := hub.Rooms[cl.ChatId]; ok {
				room := hub.Rooms[cl.ChatId]

				if _, ok := room.Clints[cl.Id]; !ok {
					room.Clints[cl.Id] = cl
				}
			}

		case cl := <-hub.Unregister:
			if _, ok := hub.Rooms[cl.ChatId]; ok {
				if _, ok := hub.Rooms[cl.ChatId].Clints[cl.Id]; ok {

					if len(hub.Rooms[cl.ChatId].Clints) != 0 {
						hub.Broadcast <- &Message{
							User_id: cl.Id,
							Chat_id: cl.ChatId,
							Content: "user left the chat",
						}
					}
					delete(hub.Rooms[cl.ChatId].Clints, cl.Id)
					close(cl.Message)
				}
			}

		case m := <-hub.Broadcast:
			if _, ok := hub.Rooms[m.Chat_id]; ok {

				for _, cl := range hub.Rooms[m.Chat_id].Clints {
					cl.Message <- m
				}
			}
		}
	}
}
