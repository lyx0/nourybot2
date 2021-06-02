// #######################################
// Just a backup file right now
// #######################################
package common

import (
	"database/sql"
	"fmt"

	"github.com/lyx0/nourybot-go/bot"
)

func JoinChannels(db *sql.DB) error {
	fmt.Println("xd")
	rows, err := db.Query("SELECT `Name` FROM `nouryqt_nourybot`.`connectchannels`")
	if err != nil {
		panic(err.Error())
	}

	// Get column names
	cols, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(cols))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// Get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var channel string
		for _, col := range values {
			if col == nil {
				channel = "NULL"
			} else {
				channel = string(col)
			}
		}
		bot.Nourybot.Client.Join(channel)
		fmt.Printf("Joined: #%s\n", channel)
	}
	return nil
}