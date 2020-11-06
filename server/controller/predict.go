package controller

import (
	"brown/model"
	"brown/service"
	"fmt"

	"github.com/kataras/iris"
)

type PredictController struct {
	Ctx     iris.Context
	Service service.PredictService
}

func NewPredictController() *PredictController {
	return &PredictController{Service: service.NewPredictServices()}
}

func (p *PredictController) Post() model.Result {
	var params map[string]string
	_ = p.Ctx.ReadJSON(&params)
	fmt.Println(params)
	result := p.Service.Predict(params["date"])
	return result
}
