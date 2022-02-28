package utils

import (
	"context"

	"github.com/andersfylling/disgord"
)

func SendResponse(session disgord.Session, handler *disgord.InteractionCreate, data *disgord.InteractionResponse) {
	session.SendInteractionResponse(
		context.Background(),
		handler,
		data,
	)
}
