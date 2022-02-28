package helpers

import (
	"github.com/andersfylling/disgord"
	"gorm.io/gorm"
)

type SlashCommand struct {
	ID                disgord.Snowflake
	Type              disgord.ApplicationCommandType
	ApplicationID     disgord.Snowflake
	GuildID           disgord.Snowflake
	Description       string
	Options           []*disgord.ApplicationCommandOption
	DefaultPermission bool
}

type Bot struct {
	Client        *disgord.Client
	SlashCommands []SlashCommand
	Database      *gorm.DB
}
