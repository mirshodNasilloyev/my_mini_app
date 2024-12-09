package minichatgo

import "time"

type TodoTweet struct {
	ID         int       `json:"id" db:"id"`
	CreatedBy  *int       `json:"created_by" db:"created_by"`
	Content    string    `json:"content" db:"content" binding:"required"`
	MediaURL   string   `json:"media_url" db:"media_url" binding:"required"`
	Created_at time.Time `json:"created_at" db:"created_at"`
}

type UserTweets struct {
	ID      int
	UserId  int
	TweetId int
}

type TodoLike struct {
	ID       int `json:"id"`
	TweetId  int `json:"tweet_id"`
	UserId   int `json:"user_id"`
	CeatedBy int `json:"created_by"`
}

type UserLikes struct {
	ID     int
	UserId int
	LikeId int
}

type TodoFollow struct {
	ID        int `json:"id"`
	UserFrom  int `json:"user_from"`
	UserTo    int `json:"user_to"`
	CreatedBy int `json:"created_by"`
}
