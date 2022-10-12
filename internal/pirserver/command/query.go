package command

import (
	"context"

	"simplepir/dal/db"
	pirpb "simplepir/grpc_gen/pir"
	"simplepir/internal/pkg/errno"
)

type QueryPiService struct {
	ctx context.Context
}

// NewQueryPiService new QueryPiService
func NewQueryPiService(ctx context.Context) *QueryPiService {
	return &QueryPiService{
		ctx: ctx,
	}
}

// QueryPi query if pi exist
func (s *QueryPiService) QueryPi(req *pirpb.PirQueryRequest) error {
	piName := req.PhoneNumber
	pis, err := db.QueryPi(s.ctx, piName)
	if err != nil {
		return err
	}
	if len(pis) == 0 {
		return errno.ErrPiNotFound
	}
	return nil
}
