package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"sync"
	"time"
)

type Conf struct {
	Addr            string
	ShutdownTimeout time.Duration
	Router          func(router *gin.Engine)
	Wg              *sync.WaitGroup
}

func InitHttp(conf Conf, middleware ...gin.HandlerFunc) {
	engine := gin.New()

	// sfsfsdsf
	engine.Use(middleware...)

	//注册路由
	conf.Router(engine)

	//监听地址
	l, err := net.Listen("tcp4", conf.Addr)
	if err != nil {
		panic("init http server listen fail :" + err.Error())
	}

	//server配置
	server := &http.Server{
		Addr:         conf.Addr,
		Handler:      engine,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		conf.Wg.Wait()
		ctx, cancel := context.WithTimeout(context.Background(), conf.ShutdownTimeout)
		defer cancel()
		err = server.Shutdown(ctx)
		fmt.Println("----- serve shutdown", time.Now().Format("2006-01-02 15:04:05.999999999"), err)
	}()

	//监听端口
	err = server.Serve(l)
	fmt.Println("----- serve close", time.Now().Format("2006-01-02 15:04:05.999999999"), err)

}
