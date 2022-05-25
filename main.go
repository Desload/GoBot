package main

import (
	"butterfly-go-bot/bot"    //we will create this later
	"butterfly-go-bot/config" //we will create this later
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
