package service

import (
	"brown/dao"
	"brown/model"
)

type PredictService interface {
	Predict(date string) (result model.Result)
}

type predictServices struct {
	dao dao.PredictDao
}

func NewPredictServices() PredictService {
	return &predictServices{
		dao: dao.NewPredictDao(),
	}
}

func (p predictServices) Predict(date string) (result model.Result) {
	//添加
	result.Data = p.dao.Predict(date)
	result.Msg = "Success"
	result.Code = 0
	return
}
