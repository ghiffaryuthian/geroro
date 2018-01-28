package user

import (
  "strings"
  "errors"
  "github.com/ghiffaryuthian/geroro/app"
  "github.com/ghiffaryuthian/geroro/log"
  "github.com/ghiffaryuthian/geroro/entity"
)

func Create(user entity.User) error {
  db    := app.Instance().Database().Begin()
  query := db.Create(&user)
  errs  := query.GetErrors()
  if len(errs) > 0 {
    db.Rollback()
    if strings.HasPrefix(errs[0].Error(), "Error 1062: Duplicate entry") {
      logger.Warn("ID:%d %s already exists in db", user.TelegramID, user.Username)
      return errors.New("You have already been registered before")
    }
    logger.Error("%v\n", errs)
  }

  logger.Info("Registered %+v\n", user)
  db.Commit()
  return nil
}
