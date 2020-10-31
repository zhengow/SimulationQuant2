package dao

import (
	"brown/datasource"
	"github.com/go-xorm/xorm"
	"brown/model"
)

type LoginDao interface {
	Login(user model.User) bool
}

type loginDao struct{
	engine *xorm.Engine
}

func NewLoginDao() LoginDao {
	return &loginDao{
		engine: datasource.InstanceMaster(),
	}
}

//登录
func (l loginDao) Login(user model.User) bool {
	has, _ := l.engine.Exist(&user)
	return has
}