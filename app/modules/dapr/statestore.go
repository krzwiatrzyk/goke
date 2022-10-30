package dapr

import (
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
)

var (
	STORE_NAME = "state"
)

func ginRegisterStateStore(g *gin.Engine) {
	g.POST("/store/:key", store)
	g.GET("/retrieve/:key", retrieve)
}

func store(c *gin.Context) {
	key := c.Param("key")
	value, _ := ioutil.ReadAll(c.Request.Body)
	if err := daprClient.SaveState(c, STORE_NAME, key, value, nil); err != nil {
		panic(err)
	}
	
}

func retrieve(c *gin.Context) {
	key := c.Param("key")
	value, err := daprClient.GetState(c, STORE_NAME, key, nil); 
	if err != nil {
		panic(err)
	}
	c.Data(http.StatusOK, "text/plain", value.Value)
}
