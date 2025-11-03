package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"mispilkabot/internal/services"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type User struct {
	UserName     string `json:"user_name"`
	IsMessaging  bool   `json:"is_messaging"`
	MessagesList []int  `json:"messages_list"`
}

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	err := services.SetSchedules(SendMessage)
	if err != nil {
		fmt.Println("BUUUUUU")
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
		addPerson(message)
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

func SendMessage(chatID string) {
	fmt.Println(chatID)
}

func getMessagesList() []int {
	return []int{0, 1, 2}
}

func addPerson(message *tgbotapi.Message) {
	data := make(map[string]User)
	raw, _ := os.ReadFile("data/users.json")

	if err := json.Unmarshal(raw, &data); err != nil {
		panic(err)
	}

	chatID := strconv.FormatInt(message.Chat.ID, 10)
	data[chatID] = User{
		UserName:     message.From.UserName,
		IsMessaging:  true,
		MessagesList: getMessagesList(),
	}

	updated, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("data/users.json", updated, 0644)
	if err != nil {
		panic(err)
	}
}