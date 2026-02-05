package commands

import "github.com/bwmarrin/discordgo"

func init() {
	cmd := &discordgo.ApplicationCommand{
		Name:        "testaddfile",
		Description: "this command run ?",
	}

	Register(cmd, testAddFileHandler)
}

func testAddFileHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Yes!",
		},
	})
	if err != nil {
		return
	}
}
