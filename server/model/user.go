package model

type User struct {
	Id         int64  `xorm:"not null pk autoincr"` //指定主键并自增
	Name       string `xorm:"unique"`               //唯一的
	Pwd        string `xorm:"not null"`
	Creat_time int64  `xorm:"created"` //创建时间
}
