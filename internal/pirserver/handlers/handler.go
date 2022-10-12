package handlers

import (
	"context"
	"fmt"

	"simplepir/dal/pack"
	pirpb "simplepir/grpc_gen/pir"
	"simplepir/internal/pirserver/command"
	"simplepir/internal/pkg/errno"
	"simplepir/internal/pkg/ttviper"
	"simplepir/third_party/forked/pir"
)

var (
	Config = ttviper.ConfigInit("PIR_SERVER", "serverConfig")
)

// CommentSrvImpl implements the last service interface defined in the IDL.
type PirSrvImpl struct {
	*pirpb.UnimplementedPirSrvServer
}

func (s *PirSrvImpl) Refresh(ctx context.Context, req *pirpb.PirRefreshRequest) (resp *pirpb.PirRefreshResponse, err error) {
	fmt.Print("2!")
	data := make([]*pir.Matrix, 1)
	hint := pir.Msg{Data: data}
	hint, err = command.NewRefreshPiService(ctx).Refresh(req)
	fmt.Print("3!")

	if err != nil {
		resp = pack.BuildpiRefreshResp(err)
		return resp, nil
	}
	resp = pack.BuildpiRefreshResp(errno.Success)
	fmt.Print("4!")

	for k, v := range hint.Data {
		resp.Data[k].Cols = v.Cols
		resp.Data[k].Rows = v.Rows
		for o, p := range v.Data {
			resp.Data[k].Data[o] = uint64(p)
		}

	}
	fmt.Print("5!")
	return resp, nil
}

func (s *PirSrvImpl) QueryPi(ctx context.Context, req *pirpb.PirQueryRequest) (resp *pirpb.PirQueryResponse, err error) {
	if req.PhoneNumber == 0 {
		resp = pack.BuildpiQueryResp(errno.ErrBind)
		return resp, nil
	}
	err = command.NewQueryPiService(ctx).QueryPi(req)

	if err != nil {
		resp = pack.BuildpiQueryResp(err)
		return resp, nil
	}
	resp = pack.BuildpiQueryResp(errno.Success)
	return resp, nil
}

// func (PirSrvImpl) mustEmbedUnimplementedPirSrvServer() {}
