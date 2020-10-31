package service

import(
	"brown/model"
	"brown/dao"
)

type RegisterService interface {
	Register(user model.User) (result model.Result)
}

type registerServices struct {
	dao dao.RegisterDao
}

func NewRegisterServices() RegisterService {
	return &registerServices{
		dao: dao.NewRegisterDao(),
	}
}

func (r registerServices) Register(user model.User) (result model.Result){
	//添加
	result.Data = r.dao.Register(user)
	result.Msg = "Success"
	result.Code = 0
	return
}