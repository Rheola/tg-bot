package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgTxt := "Here all the products: \n\n"

	products := c.productService.List()
	for _, p := range products {
		outputMsgTxt += p.Title
		outputMsgTxt += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgTxt)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", "list_10"),
		),
	)
	c.bot.Send(msg)
}
