package handler

import (
  // "strconv"
  // "time"

  // "github.com/bukalapak/geroro/entity"
  // "github.com/bukalapak/geroro/service"
  "github.com/ghiffaryuthian/geroro/app"
  "github.com/ghiffaryuthian/geroro/log"
  tb "gopkg.in/tucnak/telebot.v2"
)

func CreateUser(user *tb.User) {
  geroro := app.Instance()
  bot    := geroro.Bot()

  bot.Send(user, "hello world")
  logger.Info("Registered %s\n", user.Username)
}
