package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"notes/repository"
	"notes/service/bot"
	"os"
)

var (
	conn     *sql.DB
	host     = os.Getenv("APP_DB_HOST")
	userName = os.Getenv("APP_DB_USER")
	password = os.Getenv("APP_DB_PASSWORD")
	port     = os.Getenv("APP_DB_PORT")
)

func init() {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/telegram_bot", userName,
			password, host, port))
	if err != nil {
		panic(err)
	}
	conn = db
}

func main() {
	repo := repository.NewBotRepository(conn)

	tBot := bot.NewTelegramBot(repo)

	if err := tBot.Run(); err != nil {
		panic(err)
	}
}
