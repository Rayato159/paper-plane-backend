package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/paper-plane/configs"
	"github.com/paper-plane/internals/server"
	"github.com/paper-plane/pkg/database"
)

func main() {
	// Load dotenv config
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}
	cfg := new(configs.Config)

	// Fiber configs
	cfg.Fiber.Host = os.Getenv("SERVER_HOST")
	cfg.Fiber.Port = os.Getenv("SERVER_PORT")
	cfg.Fiber.ServerReadTimeOut = os.Getenv("SERVER_READ_TIMEOUT")

	// Database Configs
	cfg.Database.Host = os.Getenv("DB_HOST")
	cfg.Database.Port = os.Getenv("DB_PORT")
	cfg.Database.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.Database.Username = os.Getenv("DB_USERNAME")
	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	cfg.Database.Database = os.Getenv("DB_DATABASE")

	// Redis Configs
	cfg.Redis.Host = os.Getenv("REDIS_HOST")
	cfg.Redis.Port = os.Getenv("REDIS_PORT")
	cfg.Redis.Password = os.Getenv("REDIS_PASSWORD")
	cfg.Redis.DBNumber = os.Getenv("REDIS_DB_NUMBER")

	// New Database
	db, err := database.NewMySQLDBConnect(cfg)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	// Log File
	file, err := os.OpenFile(fmt.Sprintf("./pkg/log/%v.log", time.Now().Format("2006-02-01")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error, opening file: %v", err)
	}
	defer file.Close()

	s := server.NewServer(db, cfg, file)
	if err := s.Start(); err != nil {
		panic(err.Error())
	}
}
