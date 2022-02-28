package events

import (
	"counting-from-now/src/config"
	"counting-from-now/src/database/models"
	"counting-from-now/src/utils"
	"fmt"
	"log"

	"github.com/andersfylling/disgord"
)

func InteractionCreate(session disgord.Session, handler *disgord.InteractionCreate) {
	var userId disgord.Snowflake
	var user models.User

	if len(handler.Data.Options) > 0 {
		userId = disgord.ParseSnowflakeString(handler.Data.Options[0].Value.(string))
	} else {
		userId = handler.Member.UserID
	}

	disgordUser, err := session.User(userId).Get()
	if err != nil {
		log.Fatal(err)
	}

	config.Database.Find(&user, models.User{
		Id: uint(userId),
	})

	avatarUrl, _ := disgordUser.AvatarURL(1024, true)
	parsedTime := utils.ConvertTimestamp(uint(user.TimeInCall))

	utils.SendResponse(session, handler, &disgord.InteractionResponse{
		Type: 4,
		Data: &disgord.InteractionApplicationCommandCallbackData{
			Embeds: []*disgord.Embed{
				{
					Fields: []*disgord.EmbedField{
						{
							Name:  "Time in call",
							Value: parsedTime,
						},
					},
					Author: &disgord.EmbedAuthor{
						Name:    fmt.Sprintf("Information about %s", disgordUser.Username),
						IconURL: avatarUrl,
					},
				},
			},
		},
	})
}
