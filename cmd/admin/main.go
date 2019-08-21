package main

import (
	"flag"
	"github.com/Huhaokun/let-it-fail/admin"
	. "github.com/Huhaokun/let-it-fail/log"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	engine := gin.Default()

	var kubeConfig *string
	if home := homeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeConfig file")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "absolute path to the kubeConfig file")
	}
	flag.Parse()

	// use the current context in kubeConfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err.Error())
	}

	k8sClient, err := kubernetes.NewForConfig(config)

	registry := admin.NewNodeRegistry(k8sClient)

	controller := admin.NewController(registry)

	engine.GET("/api/endpoint", controller.HandleList)

	engine.POST("/api/endpoint/status_operation/:op", controller.HandleStatusOperation)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()

	err = engine.Run("0.0.0.0:7998")
	if err != nil {
		Log.Fatalf("run http server failed %v", err)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
