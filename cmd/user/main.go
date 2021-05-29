package main

import (
	"context"
	"go_python/cmd/user/cmd"
	router "go_python/cmd/user/router"
	logging "go_python/internal/pkg/logging"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	// 获取命令行参数值
	port := cmd.ArgParse().Port
	// 初始化Logger
	logging.InitWithZapLogger("user.log")
	// 初始化根router
	rootRouter := router.InitRouter()
	addr := ":" + strconv.Itoa(port)
	server := &http.Server{
		Addr:    addr,
		Handler: rootRouter,
	}

	// 退出信号接收通道
	quit := make(chan os.Signal)

	// goroutine中运行server
	go func() {
		logging.Infof("server serve and listen on %v", addr)
		err := server.ListenAndServe()

		// 服务监听失败，或者启动一个正在关闭的server时，处理错误
		if err != nil && err != http.ErrServerClosed {
			logging.Fatalf("server startup failed with err: %v", err)
		}

	}()
	// 监听退出信号
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// 接收到退出信号后处理退出逻辑
	<-quit
	logging.Warn("server is shutting down")
	err := server.Shutdown(context.Background())
	if err != nil {
		logging.Errorf("server shutdown failed with err:%v", err)
	}

}
