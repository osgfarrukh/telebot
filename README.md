# Telegram Bot Api
To install ```go get github.com/protimaru/telebot```

## Usage:
```golang
package main

import (
    "github.com/protimaru/telebot"
    "net/http"
    "log"
)

func main() {
    bot := telebot.NewBot("token")
    err := bot.SetWebhook("https://example.com")
    if err != nil {
        log.Fatalf("%+v", err)
    }
    updates := bot.WebhookUpdates()
    go http.ListenAndServe(":8000", nil)
    for update := range updates {
        if update.Message != nil {
            continue
        }
        log.Println(update.Message.Text)
    }
}

```
