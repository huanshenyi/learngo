package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const keyRequestId = "requestId"

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}

	//ミドルウェア handlersに行く前に必ずミドルウェアを通す
	r.Use(func(context *gin.Context) {
		s := time.Now()

		// ミドルウェアから出る
		context.Next()

		// path, status, log latency(待ち時間), response code logを追加
		logger.Info("incoming request",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)), //現在時間 - s
		) // {"level":"info","ts":1575721526.8082535,"caller":"gindemo/ginserver.go:20","msg":"incoming request","path":"/hello","foo":["localhost:8080"]}

	}, func(context *gin.Context) {
		// requestId をすべてのリクエストに入れる
		context.Set(keyRequestId, rand.Int())
		context.Next()
	})

	r.GET("/ping", func(context *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		// requestId 存在すれば取出す
		if rid, exists := context.Get(keyRequestId); exists {
			h[keyRequestId] = rid
		}
		context.JSON(200, h)
	})
	r.GET("/hello", func(context *gin.Context) {
		context.String(200, "hello")
	})
	r.Run()
}
