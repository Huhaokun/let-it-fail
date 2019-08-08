package log

import (
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var Log *zap.SugaredLogger

func init() {
	var err error
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("fail to init logger ", err)
	}
//	defer logger.Sync()

	Log = logger.Sugar()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		_ = <- sigs
		_ = logger.Sync()
	}()
}
