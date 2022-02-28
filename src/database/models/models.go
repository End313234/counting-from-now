package models

import (
	"time"
)

type Log struct {
	UserId   uint
	GuildId  uint
	JoinedAt time.Time
}

type User struct {
	Id         uint `gorm:"primaryKey"`
	GuildId    uint
	TimeInCall int
}
