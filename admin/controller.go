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
	NodeRegistry NodeRegistry
	Marshaller   *jsonpb.Marshaler
}

func NewController(conn *grpc.ClientConn) *Controller {
	return &Controller{
		Marshaller: &jsonpb.Marshaler{},
	}
}

func (ctrl *Controller) HandleList(c *gin.Context) {
	var allEndpoints []*contract.Endpoint

	var err error
	for _, node := range ctrl.NodeRegistry.List() {

		ctx, _ := context.WithTimeout(context.Background(), timeout)
		endpoints, err1 := node.ListEndpoint(ctx, &contract.ListCommand{})
		if err1 != nil {
			err = err1
			break
		}

		allEndpoints = append(allEndpoints, endpoints.Endpoints...)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else if jsonStr, err := ctrl.Marshaller.MarshalToString(&contract.Endpoints{Endpoints: allEndpoints}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.Data(http.StatusOK, "application/json", []byte(jsonStr))
	}
}
