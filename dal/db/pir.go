package db

import (
	"context"
	"fmt"
	"simplepir/third_party/forked/pir"

	"gorm.io/gorm"
)

// Pi Gorm Data Structures
type Pi struct {
	gorm.Model
	PhoneNumber uint64 `gorm:"index:idx_privateinfo,unique;type:varchar(256);not null" json:"privateinfo"`
}

func (Pi) TableName() string {
	return "pi"
}

// QueryPi query list of pi info
func QueryPi(ctx context.Context, piName uint64) ([]*Pi, error) {
	res := make([]*Pi, 0)
	if err := DB.WithContext(ctx).Where("pi_name = ?", piName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func Reset(ctx context.Context) (pir.Msg, error) {
	data := make([]*pir.Matrix, 1)
	msg := pir.Msg{Data: data}
	var err error
	if msg, err = InitPirDatabase(ctx); err != nil {
		return pir.Msg{}, err
	}
	return msg, nil

}

const LOGQ = uint64(32)
const SEC_PARAM = uint64(1 << 10)

var PIRDB *pir.Database

func InitPirDatabase(ctx context.Context) (pir.Msg, error) {
	N := uint64(100000)
	d := uint64(8)
	spir := pir.SimplePIR{}
	p := spir.PickParams(N, d, SEC_PARAM, LOGQ)
	var err error

	if PIRDB, err = MakePirDB(ctx, N, d, &p); err != nil {
		return pir.Msg{}, err
	}

	shared_state := spir.Init(PIRDB.Info, p)
	server_state, offline_download := spir.Setup(PIRDB, shared_state, p)

	fmt.Print(server_state, offline_download)
	return offline_download, nil
}

func MakePirDB(ctx context.Context, N, row_length uint64, p *pir.Params) (*pir.Database, error) {
	D := pir.SetupDB(N, row_length, p)
	//D.Data = pir.MatrixRand(p.l, p.m, 0, p.p)

	// Map DB elems to [-p/2; p/2]
	//D.Data.Sub(p.p / 2)

	m := make([]uint64, 100000)
	id := 0
	if err := DB.WithContext(ctx).Find(&m, "id > ?", id).Error; err != nil {
		return nil, err
	}
	for k, v := range m {
		D.Data.Data[k] = pir.Elem(v)
	}
	return D, nil
}
