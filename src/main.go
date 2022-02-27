package main

import (
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Read(".env")
	if err != nil {
		panic("could not load .env file")
	}

	client := disgord.New(disgord.Config{
		BotToken: env["BOT_TOKEN"],
	})
	defer client.Gateway().StayConnectedUntilInterrupted()

	client.Gateway().BotReady(func() {
		fmt.Println("Bot is ready to Go!")
	})
}
