package main

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/hde-oliv/ayumu/discord"
)

func main() {
	godotenv.Load() // Dev

	discord.BotToken = os.Getenv("DISCORD_TOKEN")
	discord.RunBot()
}
