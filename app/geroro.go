package app

import (
  "sync"

  "github.com/ghiffaryuthian/geroro/initializer"
  
  "github.com/jinzhu/gorm"
  tb "gopkg.in/tucnak/telebot.v2"
)

type Geroro struct {
  database *gorm.DB
  bot      *tb.Bot
}

var (
  geroro *Geroro
  once   sync.Once
)

func (g *Geroro) Database() *gorm.DB {
  return g.database
}

func (g *Geroro) Bot() *tb.Bot {
  return g.bot
}

func Instance() *Geroro {
  once.Do(func() {
    geroro = &Geroro{}
    geroro.database = initializer.Database()
    geroro.bot = initializer.GeroroBot()
  })

  return geroro
}
