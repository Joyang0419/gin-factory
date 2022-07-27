package entities

import "time"

func (UserEntity) TableName() string {
	return "sugar_parent"
}

type UserEntity struct {
	Id          int       `json:"id"`
	Vendor      string    `json:"vendor"`
	Account     string    `json:"account"`
	AccountType string    `json:"accountType"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}
