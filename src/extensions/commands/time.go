package commands

import (
	"counting-from-now/src/handlers"
	"counting-from-now/src/helpers"

	"github.com/andersfylling/disgord"
)

func init() {
	handlers.RegisterCommand("time", helpers.SlashCommand{
		ID:          0,
		Description: "The total time of the user in calls",
		Options: []*disgord.ApplicationCommandOption{
			{
				Name:        "user",
				Description: "Target of the command",
				Type:        disgord.OptionTypeUser,
				Required:    false,
			},
		},
	})
}
