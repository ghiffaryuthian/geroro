package main

import (
  "github.com/ghiffaryuthian/geroro/app"
  "github.com/ghiffaryuthian/geroro/handler"

  tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
  geroro := app.Instance()
  bot := geroro.Bot()

  bot.Handle("/register", func(m *tb.Message) {
    handler.CreateUser(m.Sender)
  })

  bot.Start()
}