package dapr

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginRegisterInvokation(g *gin.Engine) {
	g.GET("/invoke/grpc", invokeGRPC)
	g.GET("/invoke/ping", invokePing)
}

func invokeHTTP(c *gin.Context) {

}

func invokePing(c *gin.Context) {
	result, err := daprClient.InvokeMethod(c, "goke", "ping", "GET")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result))
	c.Status(http.StatusOK)
}

func invokeGRPC(c *gin.Context) {
	callGRPC()
}
