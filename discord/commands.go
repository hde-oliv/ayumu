package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "america",
		Description: "America ya!",
	},
	{
		Name:        "dollar",
		Description: "Dollar conversion rate to Brazillian Real.",
	},
}

var handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"america": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "America ya! 🇺🇸\nHALLO :D\nHALLO :D\nHALLO: D",
			},
		})
	},
	"dollar": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		rate := GetDollarRate()
		response := fmt.Sprintf("%.2f", rate)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response,
			},
		})
	},
}
