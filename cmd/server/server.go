package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"simplepir/dal"
	pirpb "simplepir/grpc_gen/pir"
	"simplepir/internal/pirserver/handlers"
	"simplepir/internal/pkg/discovery"
	"simplepir/internal/pkg/gtls"
	"simplepir/internal/pkg/ilog"
	"simplepir/internal/pkg/ttviper"
	"syscall"

	"google.golang.org/grpc"
)

var (
	Config      = ttviper.ConfigInit("PIR_SERVER", "serverConfig")
	ServiceName = Config.Viper.GetString("Server.Name")
	ServiceAddr = fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	EtcdAddress = fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	CertFile    = Config.Viper.GetString("TLS.CertFileLocalAddr")
	KeyFile     = Config.Viper.GetString("TLS.KeyFileLocalAddr")
)

func Init() {
	dal.Init()

}

func main() {
	Init()

	etcdRegister := discovery.NewRegister([]string{EtcdAddress}, ilog.New(ilog.WithLevel(ilog.InfoLevel),
		ilog.WithFormatter(&ilog.JsonFormatter{IgnoreBasicFields: false}),
	))
	defer etcdRegister.Stop()
	userNode := discovery.Server{
		Name: ServiceName,
		Addr: ServiceAddr,
	}
	if _, err := etcdRegister.Register(userNode, 10); err != nil {
		ilog.Fatalf("register server failed, err: %v", err)
	}

	tlsServer := gtls.Server{
		CertFile: CertFile,
		KeyFile:  KeyFile,
	}
	c, err := tlsServer.GetTLSCredentials()
	if err != nil {
		ilog.Fatalf("tlsServer.GetTLSCredentials err: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(c))

	pirpb.RegisterPirSrvServer(s, &handlers.PirSrvImpl{})

	lis, err := net.Listen("tcp", ServiceAddr)
	if err != nil {
		ilog.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			ilog.Fatalf("%s stopped with error: %v", ServiceName, err)
		}
	}()

	// Gracefully stop
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down pir rpc server...")

	// Delete node and revoke the given lease
	etcdRegister.Stop()

	ilog.Info("Server exiting")

}
