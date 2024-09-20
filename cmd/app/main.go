package main

import (
	"beauty_salon/internal/adapter/handler"
	"beauty_salon/internal/adapter/repository"
	"beauty_salon/internal/infrastracture/db"
	"beauty_salon/internal/infrastracture/server"
	"beauty_salon/internal/usecase"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("Failed to parse config", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Failed to load env file", err)
	}

	db, err := db.NewPostgresDB(&db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatal("failed to connect to db", err.Error())
	}

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repository)
	handler := handler.NewHandler(usecase)

	server := new(server.Server)

	log.Fatal(server.Run(viper.GetString("port"), handler.InitRouter()))
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
