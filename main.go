// main.go
package main

import (
	"encoding/json"
	"fmt"
	telegram "github.com/Syfaro/telegram-bot-api"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

const (
	BOTNAME   string = ""
	BOTAPIKEY string = ""
)

// list of actions
var (
	actions map[string]func(*gin.Context, *telegram.Update) error = map[string]func(*gin.Context, *telegram.Update) error{
		"echo":  actionEcho,
		"start": actionStart,
		"dict":  actionDictionary,
	}
	bot *telegram.BotAPI
)

func main() {
	// prepare host:port to listen to
	bind := fmt.Sprintf("%s:%s", "", "8000")
	// just predefine the err variable to avoid some problems
	var err error
	// create an instance of Telegram Bot Api
	bot, err = telegram.NewBotAPI(BOTAPIKEY)
	if err != nil {
		log.Panic(err)
	}

	// Compile the regexpression to match /action@botname
	actionSeprator := regexp.MustCompile(`^\/[\da-zA-z@]+`)

	// prepare the Gin Router
	router := gin.Default()
	// on request
	router.POST("/", func(c *gin.Context) {
		buf, err := ioutil.ReadAll(c.Request.Body)

		update := telegram.Update{}
		json.Unmarshal(buf, &update)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		// Extracting action name from text
		botName := ""
		act := actionSeprator.FindString(update.Message.Text)
		actLength := len(act)
		atsignPos := strings.Index(act, "@")
		if atsignPos != -1 {
			botName = act[atsignPos+1:]
			act = act[:atsignPos]
		}

		if botName != "" && botName != BOTNAME {
			c.String(200, "Wrong bot")
			return
		}
		act = strings.TrimPrefix(act, "/")
		act = strings.ToLower(act)
		update.Message.Text = update.Message.Text[actLength:]

		// check if the requested action exist or not
		_, has := actions[act]
		if has {
			err = actions[act](c, &update)
			if err != nil {
				c.String(500, err.Error())
				return
			}
		}

		c.String(200, "done")

	})

	router.Run(bind)
}
