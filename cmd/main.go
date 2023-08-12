package main

import (
	"log"
	"os"

	"github.com/ShatAlex/chat"
	"github.com/ShatAlex/chat/pkg/handler"
	"github.com/ShatAlex/chat/pkg/repository"
	"github.com/ShatAlex/chat/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env vars: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to initilized db: %s", err.Error())
	}

	rep := repository.NewRepository(db)
	service := service.NewService(rep)
	handlers := handler.NewHandler(service)

	server := new(chat.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
