package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"

	"GMLBot/src/commands"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://root:root@mongoDiscordBot:27017/?authSource=pass")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Impossible de contacter MongoDB :", err)
	}

	log.Println("Succès ! Connecté à MongoDB.")
	collection := client.Database("ma_base_de_donnees").Collection("utilisateurs")
	_, err = collection.InsertOne(ctx, bson.D{{"nom", "Alice"}, {"role", "admin"}})
	token := os.Getenv("TOKEN_DISCORD")
	if token == "" {
		log.Fatal("TOKEN_DISCORD is not set")
	}

	sess, err := discordgo.New("Bot " + token)

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	if err != nil {
		log.Fatalf("failed to create discord session: %v", err)
	}
	defer func(sess *discordgo.Session) {
		err := sess.Close()
		if err != nil {

		}
	}(sess)

	if err := sess.Open(); err != nil {
		log.Fatalf("failed to open discord connection: %v", err)
	}
	log.Println("Adding discordCommands...")
	sess.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			if h, ok := commands.AllHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		}
	})
	for _, v := range commands.AllCommands {
		_, err := sess.ApplicationCommandCreate(sess.State.User.ID, "", v)
		log.Printf("registered command %s", v.Name)
		if err != nil {
			log.Printf("failed to register command %s: %v", v.Name, err)
		}
	}
	log.Println("Bot is running. Press CTRL+C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	for _, v := range commands.AllCommands {
		err := sess.ApplicationCommandDelete(sess.State.User.ID, "", v.ID)
		if err != nil {

		}
	}
}
