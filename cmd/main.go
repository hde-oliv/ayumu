package main

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/hde-oliv/ayumu/discord"
	"github.com/hde-oliv/ayumu/utils"
)

func main() {
	err := godotenv.Load()
	utils.Check(err)

	discord.BotToken = os.Getenv("DISCORD_TOKEN")
	discord.RunBot()
}
