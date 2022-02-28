package main

import (
	"counting-from-now/src/config"
	"counting-from-now/src/database/models"
	"counting-from-now/src/extensions/events"
	"counting-from-now/src/helpers"
	"fmt"

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
		SlashCommands: make([]helpers.SlashCommand, 0),
	}
	defer bot.Client.Gateway().StayConnectedUntilInterrupted()

	bot.Client.Gateway().BotReady(func() {
		fmt.Println("Bot is ready to Go!")
	})

	bot.Client.Gateway().VoiceStateUpdate(events.VoiceStateUpdate)
}
