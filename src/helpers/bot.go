package helpers

import (
	"github.com/andersfylling/disgord"
	"gorm.io/gorm"
)

type SlashCommand struct {
	disgord.ApplicationCommand
}

type Bot struct {
	Client        *disgord.Client
	SlashCommands []SlashCommand
	Database      *gorm.DB
}
