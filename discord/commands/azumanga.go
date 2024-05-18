package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Bicycle3d(s *discordgo.Session, i *discordgo.InteractionCreate) {
	text := "https://vxtwitter.com/TDS_700/status/1769466287283974514"

	s.InteractionRespond(i.Interaction, textResponse(text))
}

func Maya3d(s *discordgo.Session, i *discordgo.InteractionCreate) {
	text := "https://vxtwitter.com/TDS_700/status/1766618344885506171"

	s.InteractionRespond(i.Interaction, textResponse(text))
}

func Static(s *discordgo.Session, i *discordgo.InteractionCreate) {
	text := "https://vxtwitter.com/TDS_700/status/1753908096462622773"

	s.InteractionRespond(i.Interaction, textResponse(text))
}

func Tsukurimashou3d(s *discordgo.Session, i *discordgo.InteractionCreate) {
	text := "https://vxtwitter.com/TDS_700/status/1758277384833282517"

	s.InteractionRespond(i.Interaction, textResponse(text))
}

func SataAndagi(s *discordgo.Session, i *discordgo.InteractionCreate) {
	text := "https://youtu.be/pquWcqqd_b4?feature=shared"

	s.InteractionRespond(i.Interaction, textResponse(text))
}

func Erebeta(s *discordgo.Session, i *discordgo.InteractionCreate) {
	text := "https://youtu.be/v9CzWtjK7tk"

	s.InteractionRespond(i.Interaction, textResponse(text))
}

func BicycleGta(s *discordgo.Session, i *discordgo.InteractionCreate) {
	text := "https://vxtwitter.com/OsakaDaily/status/1780798027051245951?t=tMQQNvTWIVUu3g2AtS9eIA&s=19"

	s.InteractionRespond(i.Interaction, textResponse(text))
}

func Christmas3d(s *discordgo.Session, i *discordgo.InteractionCreate) {
	text := "https://vxtwitter.com/TDS_700/status/1737957921017720879"

	s.InteractionRespond(i.Interaction, textResponse(text))
}

func Tsukurimashou(s *discordgo.Session, i *discordgo.InteractionCreate) {
	text := "https://youtu.be/_EbnC0O3_oU"

	s.InteractionRespond(i.Interaction, textResponse(text))
}

func America(s *discordgo.Session, i *discordgo.InteractionCreate) {
	text := "America ya! 🇺🇸\nHALLO :D\nHALLO :D\nHALLO: D"

	s.InteractionRespond(i.Interaction, textResponse(text))
}

func ChiyoDeath(s *discordgo.Session, i *discordgo.InteractionCreate) {
	text := "https://youtu.be/ThmZAcbjhts"

	s.InteractionRespond(i.Interaction, textResponse(text))
}
