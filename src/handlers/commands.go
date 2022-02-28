package handlers

import "counting-from-now/src/helpers"

var SlashCommands map[string]helpers.SlashCommand = make(map[string]helpers.SlashCommand)

func GetSlashCommands() map[string]helpers.SlashCommand {
	return SlashCommands
}

func RegisterCommand(name string, command helpers.SlashCommand) {
	SlashCommands[name] = command
}
