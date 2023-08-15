package repository

import (
	"fmt"
	"time"

	"github.com/ShatAlex/chat"
	"github.com/jmoiron/sqlx"
)

type ChatPostgres struct {
	db *sqlx.DB
}

func NewChatPostgres(db *sqlx.DB) *ChatPostgres {
	return &ChatPostgres{db: db}
}

func (r *ChatPostgres) Create(name string, userId int) error {

	var chatId int

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	createChat := fmt.Sprintf("INSERT INTO %s (admin_id, name) VALUES ($1, $2) RETURNING id", chatsTable)
	if err := r.db.Get(&chatId, createChat, userId, name); err != nil {
		tx.Rollback()
		return err
	}

	createChatOfUsers := fmt.Sprintf("INSERT INTO %s (chat_id, user_id) VALUES ($1, $2)", chatsOfUsersTable)
	if _, err = r.db.Exec(createChatOfUsers, chatId, userId); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *ChatPostgres) GetUserChats(userId int) ([]chat.Chat, error) {
	var chats []chat.Chat

	query := fmt.Sprintf("SELECT c.id, c.name, c.admin_id FROM %s c INNER JOIN %s cu ON c.id = cu.chat_id WHERE cu.user_id = $1", chatsTable, chatsOfUsersTable)

	err := r.db.Select(&chats, query, userId)

	return chats, err
}

func (r *ChatPostgres) GetMessages(chatId, userId int) ([]chat.Message, error) {
	var messages []chat.Message

	query := fmt.Sprintf("SELECT * FROM %s WHERE chat_id = $1 AND user_id = $2", messagesTable)

	err := r.db.Select(&messages, query, chatId, userId)

	return messages, err
}

func (r *ChatPostgres) GetUserIdByUsername(username string) (int, error) {
	var userId int

	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1", usersTable)

	err := r.db.Get(&userId, query, username)

	return userId, err
}

func (r *ChatPostgres) AddUser(chatId, userId int) error {

	query := fmt.Sprintf("INSERT INTO %s (chat_id, user_id) VALUES ($1, $2)", chatsOfUsersTable)

	_, err := r.db.Exec(query, chatId, userId)

	return err
}

func (r *ChatPostgres) CreateMessage(userId, chatId int, content string) error {

	query := fmt.Sprintf("INSERT INTO %s (user_id, chat_id, datetime, content) VALUES ($1, $2, $3, $4)", messagesTable)

	_, err := r.db.Exec(query, userId, chatId, time.Now())

	return err
}
