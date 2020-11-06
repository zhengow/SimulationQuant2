package model

type Prediction struct {
	ID    int64  `json:"id" xorm:"not null pk autoincr"` //指定主键并自增
	Stock string `json:"stock" xorm:"not null`           //唯一的
	Date  string `json:"date" xorm:"not null"`           //创建时间
}
