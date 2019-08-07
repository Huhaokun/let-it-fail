package main

import (
	"github.com/Huhaokun/let-it-fail/admin"
	. "github.com/Huhaokun/let-it-fail/log"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)


func main() {
	engine := gin.Default()

	cc, err := grpc.Dial("127.0.0.1:7999", grpc.WithInsecure())
	if err != nil {
		Log.Fatalf("dial error %v", err)
	}

	controller := admin.NewController(cc)

	engine.GET("/api/endpoint", controller.HandleList)

	err = engine.Run("0.0.0.0:7998")
	if err != nil {
		Log.Fatalf("run http server failed %v", err)
	}
}
