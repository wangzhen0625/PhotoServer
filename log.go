package main

import (
	"encoding/json"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SysLogger *zap.Logger
var err error

func GetSysLogger() *zap.Logger {
	if SysLogger == nil {
		rawJSON := []byte(`{
			"level": "debug",
			"encoding": "json",
			"outputPaths": ["stdout", "./alarmlog.txt"],
			"errorOutputPaths": ["stderr", "./alarmerr.txt"]
		  }`)
		var cfg zap.Config
		if err := json.Unmarshal(rawJSON, &cfg); err != nil {
			panic(err)
		}

		cfg.EncoderConfig = zap.NewProductionEncoderConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		SysLogger, err = cfg.Build()
		if err != nil {
			log.Fatal("init logger error: ", err)
			return nil
		}
	}
	return SysLogger
	// defer logger.Sync()
	// logger.Info("logger construction succeeded",
	// 	zap.String("url", "http://example.com"),
	// 	zap.Int("attempt", 3),
	// )

}
