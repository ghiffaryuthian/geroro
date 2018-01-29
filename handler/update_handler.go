package handler

import (
  "github.com/ghiffaryuthian/geroro/app"
  "github.com/ghiffaryuthian/geroro/log"
  tb "gopkg.in/tucnak/telebot.v2"
)

func Update(m *tb.Message) {
  geroro := app.Instance()
  bot    := geroro.Bot()
  
  btnToday := tb.InlineButton{ Unique: "today", Text: "Today's Update" }
  btnIssue := tb.InlineButton{ Unique: "issue", Text: "Issues" }
  btnTomorrow := tb.InlineButton{ Unique: "tomorrow", Text: "Tomorrow Tasks" }
  inlineKeys := [][]tb.InlineButton{
    []tb.InlineButton{btnToday},
    []tb.InlineButton{btnIssue},
    []tb.InlineButton{btnTomorrow},
  }

  if !m.Private() {
    return
  }

  bot.Send(m.Sender, "Hello!", &tb.ReplyMarkup{
    InlineKeyboard: inlineKeys,
  })

  bot.Handle(&btnToday, func (c *tb.Callback) {
    logger.Info("%+v\n", c)
    bot.Send(c.Sender, "ayayaya")
    bot.Respond(c, &tb.CallbackResponse{})
  })
}
