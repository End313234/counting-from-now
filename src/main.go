package main

import (
	"counting-from-now/src/database/models"
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

	bot := helpers.Bot{
		Client: disgord.New(disgord.Config{
			BotToken: env["BOT_TOKEN"],
		}),
		SlashCommands: []helpers.SlashCommand{},
	}
	defer bot.Client.Gateway().StayConnectedUntilInterrupted()

	bot.Client.Gateway().BotReady(func() {
		fmt.Println("Bot is ready to Go!")
	})
}
