package dapr

import (
	// "context"
	"fmt"
	"log"

	// "log"
	"net/http"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	// "github.com/dapr/go-sdk/service/common"
	// daprd "github.com/dapr/go-sdk/service/http"
	
	"github.com/gin-gonic/gin"
)

var daprClient dapr.Client


func init() {
	fmt.Println("Init function called.")
}

func Run() {
	initDaprClient, err := dapr.NewClient()
	if err != nil {
		log.Fatalf("Error initializing daprClient: %s", err)
	}
	daprClient = initDaprClient
	defer daprClient.Close()

	go runGin()

	for {
		time.Sleep(time.Second * time.Duration(5))
		fmt.Println("Sleep for 5 seconds...")
	}
}

func runGin() {
	g := gin.Default()
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	ginRegisterPubSub(g)
	ginRegisterStateStore(g)

	g.Run(fmt.Sprintf(":%d", 3000))
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
