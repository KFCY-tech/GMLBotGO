package commands

import "github.com/bwmarrin/discordgo"

var (
	AllCommands []*discordgo.ApplicationCommand
	AllHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
)

func Register(cmd *discordgo.ApplicationCommand, handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	AllCommands = append(AllCommands, cmd)
	AllHandlers[cmd.Name] = handler
}
