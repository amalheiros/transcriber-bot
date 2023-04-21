package bot

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func (app *Config) receiveMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.PostForm("MediaUrl0")

		localFile, err := app.saveAudio(file)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		//Audio from telegram and whatsapp is always on ogg format, and whisper don't acccept OGG.
		newName, err := convertToMP3(localFile)

		if err != nil {
			app.errorJSON(c, err)
			return
		}

		transcibedText, err := transcribeAudio(path.Base(newName), newName)

		if err != nil {
			app.errorJSON(c, err)
			return
		}

		reply, err := app.twilioReplyMessage(transcibedText)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		fmt.Println(reply)

		app.writeXML(c, http.StatusAccepted, reply)
	}
}

func (app *Config) callback() gin.HandlerFunc {
	return func(c *gin.Context) {
		//WIP: not implemented
	}
}

func (app *Config) sendMessage(number string, message string) gin.HandlerFunc {
	return func(c *gin.Context) {

		_, err := app.twilioSendMessage(number, message)
		if err != nil {
			app.errorJSON(c, err)
			return
		}
	}
}
