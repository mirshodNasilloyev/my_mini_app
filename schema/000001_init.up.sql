CREATE TABLE users (
    id SERIAL NOT NULL UNIQUE,
    name  VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE tweets (
    id SERIAL NOT NULL UNIQUE,
    created_by INT NULL,
    content TEXT NOT NULL,
    media_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users_tweets (
    id SERIAL NOT NULL UNIQUE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    tweet_id INT REFERENCES tweets(id) ON DELETE CASCADE NOT NULL
);