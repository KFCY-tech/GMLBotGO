package commands

import "github.com/bwmarrin/discordgo"

func init() {
	cmd := &discordgo.ApplicationCommand{
		Name:        "g",
		Description: "The team name !",
	}

	Register(cmd, gHandler)
}
func gHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "ML!",
		},
	})
	if err != nil {
		return
	}
}
