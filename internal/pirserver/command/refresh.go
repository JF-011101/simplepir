package command

import (
	"context"

	"simplepir/dal/db"
	pirpb "simplepir/grpc_gen/pir"
	"simplepir/internal/pkg/errno"
	"simplepir/third_party/forked/pir"
)

type RefreshPirService struct {
	ctx context.Context
}

func NewRefreshPiService(ctx context.Context) *RefreshPirService {
	return &RefreshPirService{
		ctx: ctx,
	}
}

func (s *RefreshPirService) Refresh(req *pirpb.PirRefreshRequest) (pir.Msg, error) {
	piPhoneNumber, err := db.Reset(s.ctx)
	if err != nil {
		return pir.Msg{}, err
	}
	if len(piPhoneNumber.Data) == 0 {
		return pir.Msg{}, errno.ErrPiNotFound
	}
	return piPhoneNumber, nil
}
