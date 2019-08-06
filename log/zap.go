package log

import (
	"go.uber.org/zap"
	"log"
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
}
