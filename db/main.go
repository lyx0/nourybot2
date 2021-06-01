package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	bot "github.com/lyx0/nourybot-go/bot"
)

// TODO: honestly idk, make it look normal i guess?
func Connect() {
	fmt.Println("Connecting to MySQL database")

	// Load .env
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")

	// Connect to MySQL database.
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+")/"+DB_NAME)
	if err != nil {
		panic(err.Error())
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Keep connection open.
	defer db.Close()
	fmt.Println("Connected to database")

	rows, err := db.Query("SELECT `Name` FROM `nouryqt_nourybot`.`connectchannels`")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var channel string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				channel = "NULL"
			} else {
				channel = string(col)
			}
			fmt.Printf(columns[i], ": ", channel, "\n")
		}
		fmt.Println("-----------------------------------")

		// TODO: move this somewhere else
		// Join each channel
		bot.Nourybot.Client.Join(channel)
		fmt.Printf("Joined: %s\n", channel)
		// Say :) in each channel
		// bot.Nourybot.Client.Say(channel, ":)")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
