package service

import (
	"brown/dao"
	"brown/model"
)

type LoginService interface {
	Login(user model.User) (result model.Result)
}

type loginServices struct {
	dao dao.LoginDao
}

func NewLoginServices() LoginService {
	return &loginServices{
		dao: dao.NewLoginDao(),
	}
}

func (l loginServices) Login(user model.User) (result model.Result) {
	//添加
	result.Data = l.dao.Login(user)
	result.Msg = "Success"
	result.Code = 0
	return
}
