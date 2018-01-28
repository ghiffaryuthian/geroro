package logger

import (
  "log"
  "os"
)
func init() {
  logfile, err := os.OpenFile("log/geroro.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
    log.Fatalf("error opening file: %v", err)
  }
  log.SetOutput(logfile)
}

func Info(format string, a ...interface{}) {
  format = "INFO: " + format
  log.Printf(format, a...)
}

func Warn(format string, a ...interface{}) {
  format = "WARN: " + format
  log.Printf(format, a...)
}

func Error(a ...interface{}) {
  log.Fatal(a...)
}