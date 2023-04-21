package bot

import "github.com/gin-gonic/gin"

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	app.Router.POST("/receive", app.receiveMessage())
	app.Router.GET("/callback", app.callback())
}
