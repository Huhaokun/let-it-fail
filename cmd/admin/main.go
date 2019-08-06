package main

import (
	. "github.com/Huhaokun/let-it-fail/log"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	err := r.Run()
	if err != nil {
		Log.Fatalf("run http server failed %v", err)
	}
}
