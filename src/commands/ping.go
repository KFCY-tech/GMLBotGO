package commands

import (
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	cmd := &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Pong",
	}

	Register(cmd, pingHandler)
}

func pingHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	mongoStatus := "MongoDB: non configuré"
	if mongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := mongoClient.Ping(ctx, nil); err != nil {
			mongoStatus = "MongoDB: erreur de connexion"
		} else {
			mongoStatus = "MongoDB: OK"
		}
	}

	content := "Pong ! \n" + mongoStatus
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
	if err != nil {
		return
	}
}
