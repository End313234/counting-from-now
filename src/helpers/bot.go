package helpers

import "github.com/andersfylling/disgord"

type SlashCommand struct {
	disgord.ApplicationCommand
}

type Bot struct {
	Client        *disgord.Client
	SlashCommands []SlashCommand
}
