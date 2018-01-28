package initializer

import (
  "os"
  "time"
  "github.com/ghiffaryuthian/geroro/log"
  "github.com/subosito/gotenv"
  tb "gopkg.in/tucnak/telebot.v2"
)

func GeroroBot() (*tb.Bot) {
  gotenv.Load()
  bot, err := tb.NewBot(tb.Settings{
    Token:  os.Getenv("BOT_TOKEN"),
    Poller: &tb.LongPoller{Timeout: 10 * time.Second},
  })
  if err == nil{
    logger.Info("GeroroBot loaded")
  } else {
    logger.Error("Failed to load GeroroBot")
  }
  return bot
}