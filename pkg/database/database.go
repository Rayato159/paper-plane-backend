package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/paper-plane/configs"
)

func NewMySQLDBConnect(cfg *configs.Config) (*sqlx.DB, error) {
	databaseUrl := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Protocol,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)
	db, err := sqlx.Connect("mysql", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, err
	}

	log.Println("MySQL Database has successfully connected. 🐬")
	return db, nil
}
