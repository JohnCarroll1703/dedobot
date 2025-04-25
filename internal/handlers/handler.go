package handlers

import (
	"dedobot/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"strings"
)

type BotHandler struct {
	bot     *tgbotapi.BotAPI
	service *service.SkufService
}

func NewBotHandler(bot *tgbotapi.BotAPI, service *service.SkufService) *BotHandler {
	return &BotHandler{bot, service}
}

func (h *BotHandler) HandleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID
		text := update.Message.Text
		var reply string
		var err error

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				reply = "Welcome to Tamagotchi Skuf Bot! Use /init to get your skuf."
			case "init":
				reply, err = h.service.InitSkuf(userID)
			case "grow":
				reply, err = h.service.FeedSkuf(userID)
			case "top":
				reply, err = h.service.ListSkufs()
			case "rename":
				args := update.Message.CommandArguments()
				if args == "" {
					reply = "Usage: /rename <new_name>"
					break
				}
				reply, err = h.service.RenameSkuf(userID, args)
			default:
				reply = "Unknown command."
			}
		} else {
			reply = h.respondToPhrase(text)
		}

		if err != nil {
			reply = "Error: " + err.Error()
		}
		if reply != "" {
			msg := tgbotapi.NewMessage(chatID, reply)
			msg.ParseMode = "Markdown"
			h.bot.Send(msg)
		}
	}
}

func (h *BotHandler) respondToPhrase(text string) string {
	lower := strings.ToLower(text)

	switch {
	case strings.Contains(lower, "привет"), strings.Contains(lower, "hi"):
		return "Я вас категорически приветствую"
	case strings.Contains(lower, "да"):
		responses := []string{
			"елда! ахаха 😂",
			"на гей пати едут поезда 🚂🌈",
		}
		return responses[rand.Intn(len(responses))]
	case strings.Contains(lower, "А"):
		return "НЕ А!"
	case strings.Contains(lower, "скуф"):
		return "В ранние 90е такой хуйни не было"
	case strings.Contains(lower, "слава дедам"):
		return "ДЕДАМ СЛАВА!!!!"
	case strings.Contains(lower, "машина"):
		return "Я позаботился о твоей машИНЕ"
	case strings.Contains(lower, "не заходи сзади"):
		return "О да, без проблем чувак."
	default:
		if rand.Intn(25) == 0 {
			return "Я полковник. Пойдем на парад?"
		}
	}
	return ""
}
