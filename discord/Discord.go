package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"github.com/Mi7teR/shavronne/commands"
	"github.com/Mi7teR/shavronne/ninja"
	"regexp"
	"strings"
)

func Run(token string) (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("error creating Discord session,", err)
		return dg, err
	}
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return dg, err
	}
	dg.UpdateStatus(0, ninja.League)
	return dg, err
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := m.Content
	if m.Author.Bot {
		return
	}
	if len(m.Mentions) > 0 {
		for _, mention := range m.Mentions {
			if mention.ID == s.State.User.ID {
				r := regexp.MustCompile("<@([0-9]+)>")
				message = strings.TrimSpace(r.ReplaceAllString(message, ""))
			}
		}
	}
	if !strings.HasPrefix(message, "!") {
		return
	}
	response := commands.Process(message)
	for _, responseMessage := range response {
		_, err := s.ChannelMessageSend(m.ChannelID, responseMessage)
		if err != nil {
			log.Println(err)
		}
	}
}
