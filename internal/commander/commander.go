package commander

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"gitlab.ozon.dev/krotovkk/homework/config"
	"gitlab.ozon.dev/krotovkk/homework/internal/handlers"
)

const (
	helpCmd   = "help"
	addCmd    = "add"
	deleteCmd = "delete"
	listCmd   = "list"
	updateCmd = "update"
)

type CmdHandler func(args string) string

type Commander struct {
	bot   *tgbotapi.BotAPI
	route map[string]CmdHandler
}

func (c *Commander) RegisterHandlers(handler *handlers.BotHandler) {
	c.registerHandler(helpCmd, handler.Help)
	c.registerHandler(addCmd, handler.AddProduct)
	c.registerHandler(deleteCmd, handler.DeleteProduct)
	c.registerHandler(listCmd, handler.ListProducts)
	c.registerHandler(updateCmd, handler.UpdateProduct)
}

func (c *Commander) registerHandler(cmd string, f func(args string) string) {
	c.route[cmd] = f
}

func (c *Commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if cmd := update.Message.Command(); cmd != "" {
			if _, ok := c.route[cmd]; ok {
				msg.Text = c.route[cmd](update.Message.CommandArguments())
			} else {
				msg.Text = fmt.Sprintf("wrong command: <%v>", cmd)
			}
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		c.bot.Send(msg)
	}

	return nil
}

func Init() (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, err
	}

	return &Commander{
		bot:   bot,
		route: map[string]CmdHandler{},
	}, nil
}
