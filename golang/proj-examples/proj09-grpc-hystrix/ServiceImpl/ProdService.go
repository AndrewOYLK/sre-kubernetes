package ServiceImpl

import (
	"context"
	"proj09-grpc-hystrix/Services"
	"strconv"
	"time"
)

func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{
		ProdID:   id,
		ProdName: pname,
	}
}

type ProdService struct {
}

func (this *ProdService) GetProdsList(ctx context.Context, in *Services.ProdsRequest, res *Services.ProdListResponse) error {
	time.Sleep(time.Second * 3)
	models := make([]*Services.ProdModel, 0)

	if in.Size == 0 {
		in.Size = 2
	}

	var i int32
	for i = 0; i < in.Size; i++ {
		models = append(models, newProd(100+i, "prodname"+strconv.Itoa(100+int(i))))
	}
	res.Data = models
	return nil
}
