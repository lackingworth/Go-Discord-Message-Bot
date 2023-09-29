package main

import (
	"github.com/lackingworth/Go-Discord-Bot/config"
	"github.com/lackingworth/Go-Discord-Bot/bot"
	"fmt"
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