package bot

import (
	"butterfly-go-bot/config" //importing our config package which we have created above
	"fmt"                     //to print errors

	"github.com/bwmarrin/discordgo" //discordgo package from the repo of bwmarrin .
)

var BotId string
var goBot *discordgo.Session

func Start() {

	//creating new bot session
	goBot, err := discordgo.New("Bot " + config.Token)

	if ErrorHandler(err) {
		return
	}

	u, err := goBot.User("@me")

	//Handlinf error

	if ErrorHandler(err) {
		return
	}

	BotId = u.ID

	goBot.AddHandler(MessageHandler)

	err = goBot.Open()

	//Error handling

	if ErrorHandler(err) {
		return
	}

	//If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")
}

//Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotId {
		return
	}

	//If we message ping to our bot in our discord it will return us pong .
	if m.Content == config.BotPrefix+"ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")

		fmt.Println(m.Author)
		fmt.Println(s.Guild("705285251227975682"))
	}
}

func ErrorHandler(err error) bool {
	fmt.Println("Error checking")

	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}
