package main

import (
	"github.com/Huhaokun/let-it-fail/admin"
	. "github.com/Huhaokun/let-it-fail/log"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)


func main() {
	engine := gin.Default()

	k8sClient, err := kubernetes.NewForConfig(&rest.Config{})

	registry := admin.NewNodeRegistry(k8sClient)

	controller := admin.NewController(registry)

	engine.GET("/api/endpoint", controller.HandleList)

	err = engine.Run("0.0.0.0:7998")
	if err != nil {
		Log.Fatalf("run http server failed %v", err)
	}
}
