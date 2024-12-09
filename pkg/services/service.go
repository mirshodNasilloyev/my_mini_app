package services

import (
	minichatgo "mini_chat_go"
	"mini_chat_go/pkg/repository"
)

type Authorization interface {
	CreateUser(user minichatgo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoTweet interface {
	CreateTweet(userId int, input minichatgo.TodoTweet) (int, error)
	GetAll(userId int) ([]minichatgo.TodoTweet, error)
	GetTweetById(userId, tweetId int) (minichatgo.TodoTweet, error)
	DeleteTweet(usetId, tweetId int) error
}
type Service struct {
	Authorization
	TodoTweet
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoTweet: NewTodoTweetService(repo.TodoTweet),
	}
}
