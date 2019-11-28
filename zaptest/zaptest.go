package main

import (
	"io"
	"log"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLogger(logpath string, loglevel string, msg string) *zap.Logger {
	hook := lumberjack.Logger{
		Filename: logpath, //log保存ファイルのpath
		MaxSize:  128,
		MaxAge:   7,
		Compress: true,
	}
	w := zapcore.AddSync(&hook)

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)

	logger := zap.New(core)
	logger.Info(msg)

	return logger
}

type Logs struct{}

func (log Logs) String() string {
	return ""
}

func main() {
	zaplog, _ := zap.NewProduction()
	zaplog.Warn("Warning test") //{"level":"warn","ts":1574900316.004626,"caller":"zaptest/zaptest.go:7","msg":"Warning test"}
	zaplog.Info("Info test")    //{"level":"info","ts":1574900570.2076318,"caller":"zaptest/zaptest.go:9","msg":"Info test"}
	logfile, err := os.OpenFile("./test.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer logfile.Close()
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("Log!!")
	initLogger("./test.log", "info", "test_info_log")

}
