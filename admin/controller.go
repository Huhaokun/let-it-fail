package admin

import (
	"context"
	"github.com/Huhaokun/let-it-fail/contract"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"time"
)

var timeout = 3 * time.Second

type Controller struct {
	ERClient contract.EndpointRegistryClient
}

func (controller *Controller) HandleList(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	endpoints, err := controller.ERClient.List(ctx, &contract.ListCommand{})
	if err != nil {
		c.JSON(500, gin.H{})
	} else {
		c.JSON(200, gin.H{})
	}
}
