package chat

type Chat struct {
	Id       int    `json:"_"`
	Name     string `json:"name" binding:"required"`
	Admin_id int    `json:"admin_id" binding:"required"`
}

type Message struct {
	Id      int    `json:"_"`
	Chat_id int    `json:"chat_id" binding:"required"`
	User_id int    `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}
