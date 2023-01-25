package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/uzbekman2005/real_time_chatting/config"
)

func ConnectToDb(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	return sqlx.Connect("postgres", psqlString)
}

func ConnectToTestDb(cfg config.Config) (*sqlx.DB, func(), error) {
	psqlString := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=business_service sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
	)

	connection, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		fmt.Println(err, "error while connection to db")
	}
	cleanUpFunc := func() {
		connection.Close()
	}
	return connection, cleanUpFunc, err
}
