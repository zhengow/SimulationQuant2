package dao

import (
	"brown/datasource"
	"brown/model"
	"fmt"

	"github.com/go-xorm/xorm"
)

type PredictDao interface {
	Predict(date string) []model.Prediction
}

type predictDao struct {
	engine *xorm.Engine
}

func NewPredictDao() PredictDao {
	return &predictDao{
		engine: datasource.InstanceMaster(),
	}
}

//登录
func (l predictDao) Predict(date string) []model.Prediction {
	fmt.Println("-----" + date)
	stocks := make([]model.Prediction, 0)
	l.engine.Where("date = ?", date).Find(&stocks)
	return stocks
}
