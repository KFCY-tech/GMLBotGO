package commands

import "github.com/bwmarrin/discordgo"
import "go.mongodb.org/mongo-driver/mongo"

var mongoClient *mongo.Client

func SetMongoClient(c *mongo.Client) {
	mongoClient = c
}

var (
	AllCommands []*discordgo.ApplicationCommand
	AllHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
)

func Register(cmd *discordgo.ApplicationCommand, handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	AllCommands = append(AllCommands, cmd)
	AllHandlers[cmd.Name] = handler
}
