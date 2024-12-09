package repository

import (
	minichatgo "mini_chat_go"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable  = "users"
	tweetsTable = "tweets"
	usersTweets = "users_tweets"
)

type Authorization interface {
	CreateUser(user minichatgo.User) (int, error)
	GetUser(username, password string) (minichatgo.User, error)
}

type TodoTweet interface {
	CreateTweet(userId int, tweet minichatgo.TodoTweet) (int, error)
	GetAll(userId int) ([]minichatgo.TodoTweet, error)
	GetTweetById(userId, tweetId int) (minichatgo.TodoTweet, error)
	DeleteTweet(usetId, tweetId int) error
}

type Repository struct {
	Authorization
	TodoTweet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoTweet:     NewTodoTweetPostgres(db),
	}
}
