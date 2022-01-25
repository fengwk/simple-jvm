package common

import "go.uber.org/zap"

// Log 通用日志组件
var Log *zap.SugaredLogger

func init() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	Log = log.Sugar()
}
