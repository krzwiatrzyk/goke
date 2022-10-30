package dapr

import "github.com/gin-gonic/gin"
import (
	"github.com/dapr/go-sdk/client"
	"io/ioutil"
)

var (
	STATE_NAME = "state"
)

func store(c *gin.Context) {
	key := nil
	value, _ := ioutil.ReadAll(c.Request.Body)
	daprClient
}

func retrieve(c *gin.Context) {

}
