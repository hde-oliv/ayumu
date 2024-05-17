package discord

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"

	"github.com/hde-oliv/ayumu/utils"
)

var BotToken string

func RunBot() {
	s, err := discordgo.New("Bot " + BotToken)
	utils.Check(err)

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	s.Open()

	for _, v := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		utils.Check(err)
	}

	defer s.Close()

	fmt.Println("Running....")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
