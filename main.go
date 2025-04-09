package dedobot

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"os"
)

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil, nil, nil)

	defer bot.StopPoll(nil,nil)
	for update := range updates {
		if update.Message
	}
}
