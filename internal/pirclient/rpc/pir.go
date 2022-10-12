package rpc

import (
	"context"
	"fmt"
	"time"

	pirpb "simplepir/grpc_gen/pir"
	"simplepir/internal/pkg/discovery"
	"simplepir/internal/pkg/errno"
	"simplepir/internal/pkg/gtls"
	"simplepir/internal/pkg/ilog"
	"simplepir/internal/pkg/ttviper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/resolver"
)

var pirClient pirpb.PirSrvClient

func InitPirRpc(Config *ttviper.Config) {
	ServerAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))

	// etcd register
	EtcdAddress := fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	EtcdResolver := discovery.NewResolver([]string{EtcdAddress}, ilog.New())
	resolver.Register(EtcdResolver)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	// init tlsClient and token
	tlsClient := gtls.Client{
		ServerName: "dytt.user.com",
		CertFile:   "D:/gopath/src/github.com/jf-011101/simplepir/config/cert/user/dytt-user.pem",
	}
	c, err := tlsClient.GetTLSCredentials()
	if err != nil {
		ilog.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}
	token := Token{
		Value: "bearer dytt.grpc.auth.token",
	}
	conn, err := RPCConnect(ctx, ServerAddress, &token, c)
	if err != nil {
		panic(err)
	}
	pirClient = pirpb.NewPirSrvClient(conn)
}

func Refresh(ctx context.Context, req *pirpb.PirRefreshRequest) (resp *pirpb.PirRefreshResponse, err error) {
	fmt.Print("1!")
	resp, err = pirClient.Refresh(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}

	return resp, nil
}

func QueryPir(ctx context.Context, req *pirpb.PirQueryRequest) (resp *pirpb.PirQueryResponse, err error) {
	resp, err = pirClient.QueryPi(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

type Token struct {
	Value string
}

const headerAuthorize string = "authorization"

// GetRequestMetadata 获取当前请求认证所需的元数据
func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{headerAuthorize: t.Value}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输
func (t *Token) RequireTransportSecurity() bool {
	return true
}

func RPCConnect(ctx context.Context, serviceAddr string, token *Token, c credentials.TransportCredentials) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(c),
		grpc.WithPerRPCCredentials(token),
	}
	conn, err = grpc.DialContext(ctx, serviceAddr, opts...)
	if err != nil {
		ilog.Fatalf("net.Connect err: %v", err)
	}
	return
}
