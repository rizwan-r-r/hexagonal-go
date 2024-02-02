package main

import (
	"context"
	"database/sql"
	"fmt"
	"hexagonalgo"
	"hexagonalgo/adapters"
	"os"
)

func main() {
	// Load our config
	conf := LoadConfig()

	// Init dependencies and pass values from config
	// - adapters
	// - parsing files
	// - setuping logger

	db, err := sql.Open("mysql", conf.MysqlDSN)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// repo := adapters.RedisRepo{} // Redis was legacy but kept still in code
	repo := &adapters.MysqlRepo{DB: db}

	app := hexagonalgo.NewApp(repo)
	app.Run(context.TODO())
}
