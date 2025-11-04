package bot

import (
	"fmt"
	"log"
	"mispilkabot/internal/services"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	err := services.SetSchedules(func(chatID string) {
		b.SendMessage(chatID)
	})

	if err != nil {
		fmt.Println("SetSchedules error: %w", err)
	}

	b.handleUpdates(b.initUpdatesChanel())
}

func (b *Bot) initUpdatesChanel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
		}
	}
}

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	switch message.Command() {
	case "start":
		services.AddPerson(message)
		msg.Text = "start command"
	case "help":
		msg.Text = "help command"
	default:
		msg.Text = "I don't know that command"
	}

	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func (b *Bot) SendMessage(chatID string) {
	data, err := services.GetPerson(chatID)
	if err != nil {
		fmt.Errorf("person data fetching error: %w", err)
		return
	}

	messagesList := data.MessagesList
	n := len(messagesList)
	last := messagesList[n-1]

	text, err := services.GetMessage(last)
	if err != nil {
		fmt.Errorf("message fetching error: %w", err)
		return
	}

	msg := tgbotapi.NewMessage(parseID(chatID), text)
	if _, err := b.bot.Send(msg); err != nil {
		log.Printf("send error to %s: %w", chatID, err)
		return
	}

	data.MessagesList = messagesList[:n-1]
	services.ChangePerson(chatID, data)
}

func parseID(s string) int64 {
	id, _ := strconv.ParseInt(s, 10, 64)
	return id
}
