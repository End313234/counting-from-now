package events

import (
	"counting-from-now/src/config"
	"counting-from-now/src/database/models"
	"time"

	"github.com/andersfylling/disgord"
)

func VoiceStateUpdate(session disgord.Session, handler *disgord.VoiceStateUpdate) {
	if handler.ChannelID == 0 {
		var log models.Log
		config.Database.Find(&log, models.Log{
			UserId:  uint(handler.UserID),
			GuildId: uint(handler.GuildID),
		})
		totalTimeInCall := time.Since(log.JoinedAt)

		var user models.User
		config.Database.Find(&user, models.User{
			Id: uint(handler.UserID),
		})

		if user.Id == 0 {
			config.Database.Create(&models.User{
				Id:         uint(handler.UserID),
				GuildId:    uint(handler.GuildID),
				TimeInCall: 0,
			})
			config.Database.Find(&user, models.User{
				Id: uint(handler.UserID),
			})
		}

		user.TimeInCall += int(totalTimeInCall.Seconds())
		config.Database.Save(&user)

		config.Database.Where("user_id = ? AND guild_id = ?", handler.UserID, handler.GuildID).Delete(&log)
	} else {
		config.Database.Create(&models.Log{
			UserId:   uint(handler.UserID),
			GuildId:  uint(handler.GuildID),
			JoinedAt: time.Now(),
		})
	}
}
