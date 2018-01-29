package handler

import (
  "github.com/ghiffaryuthian/geroro/app"
  "github.com/ghiffaryuthian/geroro/entity"
  "github.com/ghiffaryuthian/geroro/service/user"
  tb "gopkg.in/tucnak/telebot.v2"
)

func Register(message *tb.Message) {
  geroro := app.Instance()
  bot    := geroro.Bot()
  sender := message.Sender
  new_user := entity.User{ 
                TelegramID: sender.ID,
                FirstName: sender.FirstName,
                LastName: sender.LastName,
                Username: sender.Username,
              } 
  if err := user.Create(new_user); err != nil {
    bot.Reply(message, err.Error())
  } else {
    bot.Reply(message, "You have been successfully registered for daily update")
  }
}
