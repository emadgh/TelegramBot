// actions
package main

import (
	telegram "github.com/Syfaro/telegram-bot-api"
	"github.com/emadgh/vajehyab"
	"github.com/gin-gonic/gin"
)

func actionEcho(c *gin.Context, update *telegram.Update) error {
	bot.SendMessage(telegram.NewMessage(update.Message.Chat.ID, update.Message.Text))
	return nil
}

func actionStart(c *gin.Context, update *telegram.Update) error {
	bot.SendMessage(telegram.NewMessage(update.Message.Chat.ID, `Hi, I'm a Bot
You can start with the following commands

/start you just do it
/echo echo your text
/dict translate english to persian,persian to persian

Thanks for messaging me
	`))
	return nil
}

func actionDictionary(c *gin.Context, update *telegram.Update) error {
	vy := vajehyab.VajehYab{Developer: "YourDeveloperName"}
	vajeh, err := vy.Search(update.Message.Text)
	if err != nil {
		panic(err)
		return err
	}
	bot.SendMessage(telegram.NewMessage(update.Message.Chat.ID, vajeh.Data.Text.ToString()+"\nمنبع: "+vajeh.Data.Source.ToString()))
	return nil
}
