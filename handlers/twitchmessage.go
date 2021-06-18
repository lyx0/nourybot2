package handlers

import (
	"github.com/gempir/go-twitch-irc/v2"
	log "github.com/sirupsen/logrus"
)

func TwitchMessage(message twitch.PrivateMessage) {
	log.Info(message)
}