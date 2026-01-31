package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("TOKEN_DISCORD")
	if token == "" {
		log.Fatal("TOKEN_DISCORD is not set")
	}

	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("failed to create discord session: %v", err)
	}
	defer sess.Close()

	if err := sess.Open(); err != nil {
		log.Fatalf("failed to open discord connection: %v", err)
	}

	log.Println("Bot is running. Press CTRL+C to exit.")
	select {}
}
