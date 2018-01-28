package initializer

import (
  "fmt"
  "os"
  "github.com/ghiffaryuthian/geroro/log"
  "github.com/subosito/gotenv"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database() (db *gorm.DB) {
  gotenv.Load()

  dbPort := os.Getenv("DATABASE_PORT")
  if dbPort == "" {
    dbPort = "3306"
  }
  
  par := []interface{}{ os.Getenv("DATABASE_USERNAME"),
                        os.Getenv("DATABASE_PASSWORD"),
                        os.Getenv("DATABASE_HOST"),
                        string(dbPort),
                        os.Getenv("DATABASE_NAME") }

  dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", par...)
  db, err := gorm.Open("mysql", dataSourceName)
  if err == nil{
    logger.Info("Database loaded")
  } else {
    logger.Error("Failed to load Database")
  }
  return 
}
