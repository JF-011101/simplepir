package pack

import (
	"errors"

	pirpb "simplepir/grpc_gen/pir"
	"simplepir/internal/pkg/errno"
)

func BuildpiQueryResp(err error) *pirpb.PirQueryResponse {
	if err == nil {
		return piQueryResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return piQueryResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return piQueryResp(s)
}

func BuildpiRefreshResp(err error) *pirpb.PirRefreshResponse {
	if err == nil {
		return piRefreshResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return piRefreshResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return piRefreshResp(s)
}

func piQueryResp(err errno.ErrNo) *pirpb.PirQueryResponse {
	return &pirpb.PirQueryResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

func piRefreshResp(err errno.ErrNo) *pirpb.PirRefreshResponse {
	return &pirpb.PirRefreshResponse{StatusCode: int32(err.ErrCode), Data: nil, StatusMsg: &err.ErrMsg}
}
