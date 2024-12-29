package eventHandlers

import (
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	baseUtil "mrpoll_bot/base-util"
	generalModule "mrpoll_bot/general-module"
	pollModule "mrpoll_bot/poll-module"
	suggestionModule "mrpoll_bot/suggestion-module"
)

func CommandHandler(e *events.ApplicationCommandInteractionCreate) {
	commandName := e.SlashCommandInteractionData().CommandName()
	modules := []*baseUtil.Module{
		(*baseUtil.Module)(generalModule.Module),
		(*baseUtil.Module)(pollModule.Module),
		(*baseUtil.Module)(suggestionModule.Module),
	}

	for _, module := range modules {
		command, ok := module.Commands[commandName]
		if ok {
			err := command(e)
			fmt.Println("Err: ", err)

			return
		}
	}

	_ = e.CreateMessage(discord.MessageCreate{
		Content: "I couldn't find that command!",
		Flags:   discord.MessageFlagEphemeral,
	})
}
