package database

import (
	"context"
	"fmt"
	"os"

	"github.com/blastertwist/flag-dash/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg *config.Config) (*sqlx.DB, error){
	dbSource := "postgresql://" + cfg.DB.User + ":" + cfg.DB.Pass + "@" + cfg.DB.Host + ":" + cfg.DB.Port + "/" + cfg.DB.Name + "?sslmode=disable"
	conn, err := sqlx.Connect("postgres", dbSource)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	
	if err := conn.PingContext(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping the database: %v\n", err)
		os.Exit(1)
	}

	return conn, nil
}