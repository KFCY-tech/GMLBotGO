package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	token := os.Getenv("TOKEN_DISCORD")
	if token == "" {
		log.Fatal("TOKEN_DISCORD is not set")
	}

	// DiscordGo attend typiquement: "Bot <token>"
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
