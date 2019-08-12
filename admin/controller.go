package admin

import (
	"context"
	"errors"
	"github.com/Huhaokun/let-it-fail/contract"
	. "github.com/Huhaokun/let-it-fail/log"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"net/http"
	"time"
)

var timeout = 3 * time.Second

type Controller struct {
	NodeRegistry *NodeRegistry
	Marshaller   *jsonpb.Marshaler
}

func NewController(registry *NodeRegistry) *Controller {
	return &Controller{
		NodeRegistry: registry,
		Marshaller:   &jsonpb.Marshaler{},
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
	} else {
		c.JSON(http.StatusOK, &contract.Endpoints{Endpoints: allEndpoints})
	}
}

func (ctrl *Controller) HandleStatusOperation(c *gin.Context) {
	// parse request param
	operation := c.Param("op")
	endpointFilter := &contract.EndpointFilter{}
	err := c.BindJSON(endpointFilter)
	if err != nil {
		Log.Errorf("bind json error %v", err)
		return
	}

	node := ctrl.NodeRegistry.Get(endpointFilter.Host)
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	opResult := &contract.OpResult{}
	if node != nil {
		switch operation {
		case "stop":
			opResult, err = node.Stop(ctx, endpointFilter)
		case "kill":
			opResult, err = node.Kill(ctx, endpointFilter)
		case "pause":
			opResult, err = node.Pause(ctx, endpointFilter)
		default:
			err = errors.New("not support operation type")
			_ = c.AbortWithError(http.StatusNotFound, err)
			return
		}
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusOK, opResult)
	}

}
