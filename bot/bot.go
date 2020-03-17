package bot

import (
	"ajbot/config"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session


func Start() {
	goBot, err := discordgo.New("Bot "+ config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")

	// <-make(chan struct{})
	// return

}

// reponds based of chat
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		
		if m.Author.ID == BotID {
			return
		}
	
		if m.Content == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		}

		if m.Content == "!hiphip" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "hurray")
		}
		
	}
	fmt.Println(m.Content)
}

	




