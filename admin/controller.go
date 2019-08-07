package admin

import (
	"context"
	"github.com/Huhaokun/let-it-fail/contract"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

var timeout = 3 * time.Second

type Controller struct {
	ERClient   contract.EndpointRegistryClient
	Marshaller *jsonpb.Marshaler
}

func NewController(conn *grpc.ClientConn) *Controller {
	return &Controller{
		ERClient:   contract.NewEndpointRegistryClient(conn),
		Marshaller: &jsonpb.Marshaler{},
	}
}

func (ctrl *Controller) HandleList(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	endpoints, err := ctrl.ERClient.List(ctx, &contract.ListCommand{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else if jsonStr, err := ctrl.Marshaller.MarshalToString(endpoints); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.Data(http.StatusOK, "application/json", []byte(jsonStr))
	}
}
