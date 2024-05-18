package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/hde-oliv/ayumu/discord/commands"
)

var description = []*discordgo.ApplicationCommand{
	{
		Name:        "america",
		Description: "America ya!",
	},
	{
		Name:        "bicycle-3d",
		Description: "Chiyo and Osaka on a bicycle, 3D version.",
	},
	{
		Name:        "tsukurimashou",
		Description: "Chiyo cooking.",
	},
	{
		Name:        "christmas-3d",
		Description: "Osaka saying merry christmas, 3D version.",
	},
	{
		Name:        "bicycle-gta",
		Description: "Chiyo and Osaka on a bicycle, GTA version.",
	},
	{
		Name:        "erebeta",
		Description: "Chiyo and Osaka in an elevator.",
	},
	{
		Name:        "sata-andagi",
		Description: "Osaka saying 'Sata andagi'.",
	},
	{
		Name:        "tsukurimashou-3d",
		Description: "Chiyo cooking, 3D version.",
	},
	{
		Name:        "static",
		Description: "Osaka playing with Chiyo's hair.",
	},
	{
		Name:        "maya-3d",
		Description: "Sakaki and Maya, 3D version",
	},
	{
		Name:        "chiyo-death",
		Description: "RIP Chiyo-chan.",
	},
	{
		Name:        "dollar",
		Description: "Dollar conversion rate to Brazillian Real.",
	},
}

var handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"america":          commands.America,
	"bicycle-3d":       commands.Bicycle3d,
	"tsukurimashou":    commands.Tsukurimashou,
	"christmas-3d":     commands.Christmas3d,
	"bicycle-gta":      commands.BicycleGta,
	"erebeta":          commands.Erebeta,
	"sata-andagi":      commands.SataAndagi,
	"tsukurimashou-3d": commands.Tsukurimashou3d,
	"static":           commands.Static,
	"maya-3d":          commands.Maya3d,
	"chiyo-death":      commands.ChiyoDeath,
	"dollar":           commands.Dollar,
}
