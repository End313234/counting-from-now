package main

import (
	"counting-from-now/src/config"
	"counting-from-now/src/database/models"
	_ "counting-from-now/src/extensions/commands"
	"counting-from-now/src/extensions/events"
	"counting-from-now/src/handlers"
	"counting-from-now/src/helpers"
	"fmt"
	"log"

	"github.com/andersfylling/disgord"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	env, err := godotenv.Read(".env")
	if err != nil {
		panic("could not load .env file")
	}

	database, err := gorm.Open(sqlite.Open("database.sqlite"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.User{}, &models.Log{})

	config.Database = database
	bot := helpers.Bot{
		Client: disgord.New(disgord.Config{
			BotToken: env["BOT_TOKEN"],
			Intents:  disgord.AllIntents(),
		}),
	}
	defer bot.Client.Gateway().StayConnectedUntilInterrupted()

	bot.Client.Gateway().BotReady(func() {
		for name, command := range handlers.GetSlashCommands() {
			slashCommand := &disgord.CreateApplicationCommand{
				Name:              name,
				Description:       command.Description,
				Type:              command.Type,
				Options:           command.Options,
				DefaultPermission: command.DefaultPermission,
			}

			if err := bot.Client.ApplicationCommand(command.ID).Guild(907026890546622555).Create(slashCommand); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("Bot is ready to Go!")
	})

	bot.Client.Gateway().VoiceStateUpdate(events.VoiceStateUpdate)
	bot.Client.Gateway().InteractionCreate(events.InteractionCreate)
}
