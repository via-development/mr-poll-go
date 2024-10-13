package mpHandlers

import (
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	generalModule "mrpoll_go/general-module"
	pollModule "mrpoll_go/poll-module"
)

func CommandHandler(e *events.ApplicationCommandInteractionCreate) {
	commandName := e.SlashCommandInteractionData().CommandName()

	command, ok := generalModule.Module.Commands[commandName]
	if ok {
		err := command(e)
		fmt.Println(err)

		return
	}

	command, ok = pollModule.Module.Commands[commandName]
	if ok {
		err := command(e)
		fmt.Println(err)

		return
	}

	_ = e.CreateMessage(discord.MessageCreate{
		Content: "I couldn't find that command!",
		Flags:   discord.MessageFlagEphemeral,
	})
}
