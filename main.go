package main

import (
  "context"
  "os"
  "os/signal"
  "syscall"

  "github.com/go-telegram/bot"
  "go.uber.org/zap"

  "github.com/HappyDevopsClub/tg-chat-bot/handlers"
)

var log = zap.L()

func main() {
  ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
  defer cancel()

  var opts []bot.Option

  if os.Getenv("DEBUG") == "true" {
    opts = append(opts, bot.WithDebug())
  }

  b, err := bot.New(
    os.Getenv("TELEGRAM_BOT_TOKEN"),
    opts...,
  )
  if err != nil {
    log.Fatal("failed to create bot", zap.Error(err))
  }

  b.RegisterHandlerMatchFunc(handlers.WelcomeMessageFilter, handlers.WelcomeMessageHandler)

  b.Start(ctx)
}
