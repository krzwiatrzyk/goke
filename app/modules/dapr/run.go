package dapr

import (
	// "context"
	"fmt"
	// "log"
	"net/http"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	// "github.com/dapr/go-sdk/service/common"
	// daprd "github.com/dapr/go-sdk/service/http"
	"github.com/gin-gonic/gin"
)

var daprClient dapr.Client

var (
	PUBSUB_NAME = "queue"
	PUBSUB_TOPIC = "goke"
)

func Run() {
	daprClient, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer daprClient.Close()

    go runGin()

	for {
		time.Sleep(time.Second * time.Duration(5))
		fmt.Println("Sleep for 5 seconds...")
	}
}

func runGin() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/queue/publish", publish)
	r.GET("/queue/subscribe", subscribe)
	r.GET("/dapr/subscribe", informDaprSubscribe)

	r.Run(fmt.Sprintf(":%d", 3000))
}

func publish(c *gin.Context) {
	if err := daprClient.PublishEvent(c, PUBSUB_NAME, PUBSUB_TOPIC, c.Request.Body); err != nil {
		c.Status(http.StatusInternalServerError)
		panic(err)
	}

	c.Status(http.StatusOK)
}

func informDaprSubscribe(c *gin.Context) {
  response := []map[string]interface{}{
	{
		"pubsubname": PUBSUB_NAME,
		"topic": PUBSUB_TOPIC,
		"route": "/queue/subscribe",
		"metadata": map[string]string{
			"rawPayload": "true",
		},
	},
  }

  c.JSON(http.StatusOK,response)
}

func subscribe(c *gin.Context) {
	fmt.Println(c.Request.Body)
	c.Status(http.StatusOK)
}

// func runDapr() {
// 	sub := &common.Subscription{
// 		PubsubName: PUBSUB_NAME,
// 		Topic:      PUBSUB_TOPIC,
// 		Route:      "/events",
// 	}


// 	s := daprd.NewService(":" + "3001")

// 	err := s.AddTopicEventHandler(sub, eventHandler)
// 	if err != nil {
// 		log.Fatalf("error adding topic subscription: %v", err)
// 	}

// 	if err := s.Start(); err != nil && err != http.ErrServerClosed {
// 		log.Fatalf("error: %v", err)
// 	}
// }

// func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
// 	log.Printf("event - PubsubName:%s, Topic:%s, ID:%s, Data: %v", e.PubsubName, e.Topic, e.ID, e.Data)
// 	return true, nil
// }