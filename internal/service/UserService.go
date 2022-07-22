package service

import (
	"aCupOfGin/internal/dao"
	"aCupOfGin/internal/encryptService"
	"aCupOfGin/internal/entity"
)

func CreateUser(user *entity.User) (err error) {
	err, user.Password = encryptService.HashAndSalt(user.Password)
	if err != nil {
		return err
	}

	if err = dao.SqlSession.Select(
		"Vendor",
		"Account", "AccountType", "Password", "Name").Create(user).Error; err != nil {
		return err
	}
	return
}

func GetAllUser() (userList []*entity.User, err error) {
	if err := dao.SqlSession.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func DeleteUserById(id string) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.User{}).Error
	return
}

func GetUserById(id string) (user *entity.User, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateUser(user *entity.User, isPwdUpdated bool) (err error) {
	if isPwdUpdated {
		err, user.Password = encryptService.HashAndSalt(user.Password)
		if err != nil {
			return err
		}
	}

	err = dao.SqlSession.Save(user).Error
	return
}
