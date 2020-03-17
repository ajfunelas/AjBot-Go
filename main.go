package main

import (
	"fmt"
	"ajbot/config"
	"ajbot/bot"

	"ajbot/funcs"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()


	rss, err := funcs.ReadRSSFeed("https://www.reddit.com/r/CoronavirusDownunder.rss")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(rss)

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