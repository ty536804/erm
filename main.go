package main

import (
	"context"
	_ "erm/Commands"
	"erm/Pkg/Setting"
	"erm/Router"
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := Router.InitRouter()
	ginpprof.Wrap(router)
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", Setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    Setting.ReadTimeout,
		WriteTimeout:   Setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Printf("Actual pid is %d", syscall.Getpid())
}
