package bot

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type jsonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (app *Config) errorJSON(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	c.JSON(statusCode, jsonResponse{Status: statusCode, Message: err.Error()})
}

func (app *Config) writeJSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, jsonResponse{Status: status, Message: "success", Data: data})
}

func (app *Config) writeXML(c *gin.Context, status int, data string) {
	c.Header("Content-Type", "text/xml")
	c.String(status, data)
}

func (app *Config) saveAudio(fileName string) (string, error) {
	id := uuid.New()

	localFileName := "media/" + id.String() + ".ogg"

	out, err := os.Create(localFileName)
	defer out.Close()

	resp, err := http.Get(fileName)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return localFileName, nil
}
