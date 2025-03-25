package main

import (
	"fmt"
	_ "github.com/4epuha1337/botick/db"
	"github.com/4epuha1337/botick/tools"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strings"
)

var IntroStr = `Здравствуйте, я бот, который <вставить, можно несколько строк>`
var AddReqStr = `Опишите вашу проблему.`

var IsAddReq = false

var IdAdm = os.Getenv("TELEGRAM_IDADM")

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if IsAddReq {
			text := update.Message.Text
			
			continue
		}
		
		if update.Message == nil {
			continue
		}

		text := update.Message.Text

		//fromID := fmt.Sprintf("%d", update.Message.From.ID)
		if strings.HasPrefix(text, "/start") {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, IntroStr)
			if _, err := bot.Send(msg); err != nil {
				fmt.Printf("Error sending message: %v\n", err)
				errorMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, there was an error processing your message")
				bot.Send(errorMsg)
				continue
			}
		}

		if strings.HasPrefix(text, "/requests") {
			if tools.IsAdmin(update.Message.From.ID) {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Current requests:\n")
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
				continue
			}
		}

	if strings.HasPrefix(text, "/newrequest") {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, AddReqStr)
		if _, err := bot.Send(msg); err != nil {
			fmt.Printf("Error sending message: %v\n", err)
			errorMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, there was an error processing your message")
			bot.Send(errorMsg)
			continue
		}
		IsAddReq = true
	}
		
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			fmt.Printf("Error sending message: %v\n", err)
			errorMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, there was an error processing your message")
			bot.Send(errorMsg)
			continue
		}

	}
}

// /requests - вывод списка поступивших запросов (для админа)
// добавить запись запроса (пока нет тз - одна строка)
// бд для записи запросов (хз, пока sqlite можно, потом postgres хуйнуть)
// мб админid в бд, чтобы несколько админов могли получать запросы (сделал через перменные окружения+недопарсер) +
// защитить admid?? (если в бд вытаскиваю) -
// команды для юзеров??? (как будто просто чата хватит, но тз опять же) +
//
