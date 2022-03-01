package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"time"
)

func InitHttp(addr string, ctx context.Context, registerRouter func(router *gin.Engine), middleware ...gin.HandlerFunc) {
	engine := gin.New()

	// sfsfsdsf
	engine.Use(middleware...)

	//注册路由
	registerRouter(engine)

	//监听地址
	l, err := net.Listen("tcp4", addr)
	if err != nil {
		panic("init http server listen fail :" + err.Error())
	}

	//server配置
	server := &http.Server{
		Addr:         addr,
		Handler:      engine,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		//监听端口
		_ = server.Serve(l)
	}()

	_ = server.Shutdown(ctx)
}
