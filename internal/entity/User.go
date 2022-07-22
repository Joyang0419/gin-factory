package entity

import "time"

func (User) TableName() string {
	return "sugar_parent"
}

type User struct {
	Id          uint64    `json:"id"`
	Vendor      string    `json:"vendor"`
	Account     string    `json:"account"`
	AccountType string    `json:"accountType"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	CreateTime  time.Time `json:"createTime"`
}
