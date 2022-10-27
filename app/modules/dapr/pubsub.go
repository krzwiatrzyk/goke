package dapr

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/gin-gonic/gin"
)

var (
	PUBSUB_NAME  = "queue"
	PUBSUB_TOPIC = "goke"
)

var messages []string

func ginRegisterPubSub(g *gin.Engine) {
	g.POST("/queue/publish", publish)
	g.POST("/queue/subscribe", subscribe)
	g.GET("/dapr/subscribe", informDaprSubscribe)
	g.GET("/queue/list", list)
}

func publish(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := daprClient.PublishEvent(c, PUBSUB_NAME, PUBSUB_TOPIC, string(body)); err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	c.Status(http.StatusOK)
}

// Optional information regarding subscription
func informDaprSubscribe(c *gin.Context) {
	response := []map[string]interface{}{
		{
			"pubsubname": PUBSUB_NAME,
			"topic":      PUBSUB_TOPIC,
			"route":      "/queue/subscribe",
			"metadata": map[string]string{
				"rawPayload": "false",
			},
		},
	}

	c.JSON(http.StatusOK, response)
}

func subscribe(c *gin.Context) {
	event, _ := cloudevents.NewEventFromHTTPRequest(c.Request)
	messages = append(messages, string(event.Data()))
	c.Status(http.StatusOK)
}

func list(c *gin.Context) {
	response, _ := json.Marshal(messages)
	c.Data(http.StatusOK, "", response)
}