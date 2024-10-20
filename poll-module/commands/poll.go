package pollCommands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"mrpoll_go/assets"
	pollAssets "mrpoll_go/poll-module/assets"
)

func PollCommand(interaction *events.ApplicationCommandInteractionCreate) error {
	subcommand := interaction.SlashCommandInteractionData().SubCommandName
	if subcommand == nil {
		return nil
	}

	switch *subcommand {
	case "yes-or-no", "multiple-choice", "single-choice":
		return pollCreateCommand(interaction)
	case "list":
		return pollListCommand(interaction)
	case "end":
		return pollEndCommand(interaction)
	case "online":
		return pollOnlineCommand(interaction)
	default:
		return interaction.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				assets.MakeSimpleEmbed("I cannot find that command!"),
			},
		})
	}
}

func pollCreateCommand(interaction *events.ApplicationCommandInteractionCreate) error {
	return interaction.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			pollAssets.MakePollEmbed(),
		},
	})
}

func pollOnlineCommand(interaction *events.ApplicationCommandInteractionCreate) error {
	return interaction.CreateMessage(discord.MessageCreate{
		Components: []discord.ContainerComponent{
			discord.ActionRowComponent{
				discord.ButtonComponent{ // Leaked?!?!
					URL:   "https://mrpoll.dev/polls",
					Style: discord.ButtonStyleLink,
					Label: "Create Poll",
				},
			},
		},
	})
}

func pollEndCommand(interaction *events.ApplicationCommandInteractionCreate) error {
	return nil
}

func pollListCommand(interaction *events.ApplicationCommandInteractionCreate) error {
	return nil
}
