package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(message *tgbotapi.Message) {
	args := message.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args, err)
	}

	product, err := c.productService.Get(arg)
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", args, err)
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		product.Title,
	)
	c.bot.Send(msg)
}
