package repository

import (
	"fmt"
	minichatgo "mini_chat_go"

	"github.com/jmoiron/sqlx"
)

type TodoTweetPostgres struct {
	db *sqlx.DB
}

func NewTodoTweetPostgres(db *sqlx.DB) *TodoTweetPostgres{
	return &TodoTweetPostgres{db: db}
}

func (ttp *TodoTweetPostgres) CreateTweet(userId int, tweet minichatgo.TodoTweet) (int, error) {
	tx, err := ttp.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createTweetQuery := fmt.Sprintf("INSERT INTO %s (content, media_url) VALUES ($1, $2) RETURNING id", tweetsTable)
	row := tx.QueryRow(createTweetQuery, tweet.Content, tweet.MediaURL)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	fmt.Print(id)
	createUsersTweetQuery := fmt.Sprintf("INSERT INTO %s (user_id, tweet_id) VALUES ($1, $2)", usersTweets)
	_, err = tx.Exec(createUsersTweetQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("Error committing transaction:", err)
		return 0, err
	}
	return id, nil
}

func (ttp *TodoTweetPostgres) GetAll(userId int) ([]minichatgo.TodoTweet, error) {
	var tweets []minichatgo.TodoTweet

	query := fmt.Sprintf("SELECT tt.id, tt.created_by, tt.content, tt.media_url, tt.created_at FROM %s tt INNER JOIN %s ut on tt.id = ut.tweet_id WHERE ut.user_id = $1", tweetsTable, usersTweets)
	err := ttp.db.Select(&tweets, query, userId)
	return tweets, err
}

func (ttp *TodoTweetPostgres) GetTweetById(userId, tweetId int) (minichatgo.TodoTweet, error) {
	var tweet minichatgo.TodoTweet

	query := fmt.Sprintf("SELECT tt.id, tt.created_by, tt.content, tt.media_url, tt.created_at FROM %s tt INNER JOIN %s ut on tt.id = ut.tweet_id WHERE ut.user_id = $1 AND ut.tweet_id = $2", tweetsTable, usersTweets)
	err := ttp.db.Get(&tweet, query, userId, tweetId)
	return tweet, err
}

func (ttp *TodoTweetPostgres) DeleteTweet(userId, tweetId int) error {

	query := fmt.Sprintf("DELETE FROM %s tt USING %s ut WHERE tt.id = ut.tweet_id AND ut.user_id=$1 AND ut.tweet_id=$2", tweetsTable, usersTweets)
	_, err := ttp.db.Exec(query, userId, tweetId)
	return err
}