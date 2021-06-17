package main

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/lyx0/nourybot2/client"
	"github.com/lyx0/nourybot2/config"
	"github.com/lyx0/nourybot2/db"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Load config containing client configurations.
	cfg := config.LoadConfig()

	// Create a new twitch client with parameters specified from
	// the config module.
	twitchClient := twitch.NewClient(cfg.Username, cfg.Oauth)

	// Create a new sql connection.
	sqlClient := db.Connect(cfg)
	defer sqlClient.Close()

	// Create New Bot with twitch client and
	// config and connect to chat.
	twitchBot := client.NewTwitchBot(cfg, twitchClient, sqlClient)

	// Connect to Discord
	go func() {
		client.DiscordConnect()
	}()

	// Connect to Twitch
	err := twitchBot.Connect()
	if err != nil {
		log.Fatal("Couldn't establish Twitch connection", err)
	}
}
