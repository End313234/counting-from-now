package models

import (
	"time"
)

type Log struct {
	UserId   uint `gorm:"primaryKey"`
	GuildId  uint
	JoinedAt time.Time
}

type User struct {
	UserId     uint `gorm:"primaryKey"`
	GuildId    uint
	TimeInCall time.Time
}
