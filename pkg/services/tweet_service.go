package services

import (
	minichatgo "mini_chat_go"
	"mini_chat_go/pkg/repository"
)

type TodoTweetService struct {
	repo repository.TodoTweet
}

func NewTodoTweetService(repo repository.TodoTweet) *TodoTweetService {
	return &TodoTweetService{repo: repo}
}

func (tts *TodoTweetService) CreateTweet(userId int, input minichatgo.TodoTweet) (int, error) {
	return tts.repo.CreateTweet(userId, input)
}

func (tts *TodoTweetService) GetAll(userId int) ([]minichatgo.TodoTweet, error) {
	return tts.repo.GetAll(userId)
}

func (tts *TodoTweetService) GetTweetById(userId, tweetId int) (minichatgo.TodoTweet, error) {
	return tts.repo.GetTweetById(userId, tweetId)
}

func (tts *TodoTweetService) DeleteTweet(userId, tweetId int) error {
	return tts.DeleteTweet(userId, tweetId)
}
