package dao

import (
	"brown/datasource"
	"github.com/go-xorm/xorm"
	"brown/model"
)

type RegisterDao interface {
	Register(user model.User) (model.User)
}

type registerDao struct{
	engine *xorm.Engine
}

func NewRegisterDao() RegisterDao {
	return &registerDao{
		engine: datasource.InstanceMaster(),
	}
}

//登录
func (r registerDao) Register(user model.User) (model.User) {
	r.engine.Insert(user)
	return user
}