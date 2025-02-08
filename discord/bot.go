package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"

	"github.com/hde-oliv/ayumu/utils"
)

var BotToken string

func RunBot() {
	s, err := discordgo.New("Bot " + BotToken)
	utils.Check(err)

	defer s.Close()

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	if err := s.Open(); err != nil {
		log.Fatal(err)
	}

	for _, c := range description {
		for _, g := range s.State.Guilds {
			_, err := s.ApplicationCommandCreate(s.State.User.ID, g.ID, c)
			fmt.Printf("Added command %s to %s\n", c.Name, g.Name)

			utils.Check(err)
		}
	}

	fmt.Println("Running....")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
