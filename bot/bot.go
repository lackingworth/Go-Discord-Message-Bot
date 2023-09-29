package bot

import (
	"github.com/lackingworth/Go-Discord-Bot/config"
	"github.com/bwmarrin/discordgo"
	"fmt"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return 
	}

	if m.Content == "HAI" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "KTHXBAI")
	}
}