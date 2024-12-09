package main

import (
	"fmt"
	minichatgo "mini_chat_go"
	"mini_chat_go/pkg/handlers"
	"mini_chat_go/pkg/repository"
	"mini_chat_go/pkg/services"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config %s", err.Error())
	}
	fmt.Print("Config file initialized\n")

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error get env file %s", err.Error())
	}
	fmt.Print("Env file initialized\n")

	db, err := minichatgo.NewPostgresDB(minichatgo.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	fmt.Print("DB initialized\n")

	if err != nil {
		logrus.Fatalf("error connecting db %s", err.Error())
	}
	repository := repository.NewRepository(db)
	service := services.NewService(repository)
	handlers := handlers.NewHandler(service)
	srv := new(minichatgo.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitHandlers()); err != nil {
		logrus.Fatalf("failed to start server %s", err.Error())
	}
	logrus.Print("mini_chat_go app running....")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
