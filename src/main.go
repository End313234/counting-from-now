package main

import (
	"counting-from-now/src/helpers"
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Read(".env")
	if err != nil {
		panic("could not load .env file")
	}

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
