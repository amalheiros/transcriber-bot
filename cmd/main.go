package main

import (
	bot "transcriber-bot/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	app := bot.Config{Router: router}

	app.Routes()

	router.Run(":8080")
}
