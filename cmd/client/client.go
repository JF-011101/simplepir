package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"simplepir/internal/pirclient/handlers"
	"simplepir/internal/pirclient/rpc"

	"simplepir/internal/pkg/ilog"
	"simplepir/internal/pkg/ttviper"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

var (
	Config      = ttviper.ConfigInit("PIR_CLIENT", "clientConfig")
	ServiceName = Config.Viper.GetString("Server.Name")
	ServiceAddr = fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
)

func main() {
	rpc.InitPirRpc(&Config)

	r := gin.New()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, false))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	pir := r.Group("/pir")

	pir.POST("/refresh/", handlers.Refresh)
	pir.GET("/query/", handlers.QueryPirBoundary)
	pir.POST("/query/", handlers.QueryPir)

	srv := &http.Server{
		Addr:    ":8088",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ilog.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ilog.Info("Shutting down http server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		ilog.Fatal("Server forced to shutdown:", err)
	}

	ilog.Info("Server exiting")
}
