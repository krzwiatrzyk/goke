package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var rainHTML []byte

const (
    ContentTypeBinary = "application/octet-stream"
    ContentTypeForm   = "application/x-www-form-urlencoded"
    ContentTypeJSON   = "application/json"
    ContentTypeHTML   = "text/html; charset=utf-8"
    ContentTypeText   = "text/plain; charset=utf-8"
)

func Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/headers", returnHeaders)
	r.GET("/rain", rain)

	r.Run(fmt.Sprintf(":%d", 3000))
}

func returnHeaders(c *gin.Context) {
	c.JSON(http.StatusOK, c.Request.Header)
}

func rain(c *gin.Context) {
	c.Data(http.StatusOK, ContentTypeHTML, rainHTML)
}

func init() {
  html, err := ioutil.ReadFile("static/rain/index.html")
  rainHTML = html

  if err != nil {
	panic(err)
  }
}

