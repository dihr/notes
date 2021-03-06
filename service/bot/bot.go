package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"notes/repository"
	"os"
	"strings"
)

type (
	TelegramBot interface {
		Run() error
	}

	telegramBotImp struct {
		db repository.BotRepository
	}
)

var (
	botToken = os.Getenv("TELEGRAM_BOT_TOKEN")
)

func NewTelegramBot(db repository.BotRepository) TelegramBot {
	return &telegramBotImp{
		db: db,
	}
}
func (t *telegramBotImp) Run() error {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "help", "categories", "start":
			msg.Text, err = t.callForHelpOrCategories()
			if err != nil {
				msg.Text = err.Error()
			}
			if _, err := bot.Send(msg); err != nil {
				return err
			}
		default:
			var hasImage bool
			msg.Text, hasImage, err = t.callForSubCategoryOrCategory(update.Message.Command())
			if err != nil {
				msg.Text = err.Error()
			}
			if _, err := bot.Send(msg); err != nil {
				return err
			}
			if hasImage {
				photoBytes, err := ioutil.ReadFile(fmt.Sprintf("/images/%s.png",
					update.Message.Command()))
				if err != nil {
					return err
				}
				photoFileBytes := tgbotapi.FileBytes{
					Name:  "picture",
					Bytes: photoBytes,
				}
				if _, err := bot.Send(tgbotapi.NewPhoto(msg.ChatID, photoFileBytes)); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (t *telegramBotImp) callForHelpOrCategories() (string, error) {
	categories, err := t.db.GetAllCategories()
	if err != nil {
		return "", err
	}
	var message string
	for _, c := range categories {
		message += fmt.Sprintf("/%s\n", c.Name)
	}
	return message, nil
}

func (t *telegramBotImp) callForSubCategoryOrCategory(command string) (string, bool, error) {
	var message string
	var hasFile bool
	categories, err := t.db.GetAllCategories()
	if err != nil {
		return "", false, err
	}
	for _, c := range categories {
		if command == c.Name {
			subCategories, err := t.db.GetSubCategories("category_id = ?", c.ID)
			if err != nil {
				return "", false, err
			}
			for _, sc := range subCategories {
				message += fmt.Sprintf("/%s\n", sc.Name)
			}
		}
	}
	subCategories, err := t.db.GetAllSubCategories()
	if err != nil {
		return "", false, err
	}
	for _, sub := range subCategories {
		if sub.Name == command {
			hasFile = sub.FlagImg
			// Use pipes to separate new lines in text from database.
			messageSplit := strings.Split(sub.Text, "|")
			for _, m := range messageSplit {
				message += fmt.Sprintf("%s\n", m)
			}
		}
	}
	return message, hasFile, nil
}
