package main

import (
	"fmt"
	"./config"
	"./bot"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}




// func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	
// 	if m.Author.ID == BotID {
// 		return
// 	}

// 	if m.Content == "ping" {
// 		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
// 	}
	
// 	// fmt.Println(m.Content)
// }