package userRepo

import (
	"aCupOfGin/internal/entities"
	"aCupOfGin/internal/tools/dbManager"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type GormUserRepo struct {
	DBManager dbManager.InterfaceDBManger
}

func (userRepo *GormUserRepo) CreateUser(
	vendor string,
	account string,
	accountType string,
	hashedPassword string,
	name string) (bool, error) {

	user := entities.UserEntity{
		Vendor:      vendor,
		Account:     account,
		AccountType: accountType,
		Password:    hashedPassword,
		Name:        name,
	}

	DBConnection := userRepo.DBManager.ProvideDBConnection().(gorm.DB)

	result := DBConnection.Select("Vendor", "Account", "AccountType", "Password", "Name").Create(
		&user)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil

}

func (userRepo *GormUserRepo) GetAllUsers() []entities.UserEntity {
	var users []entities.UserEntity

	DBConnection := userRepo.DBManager.ProvideDBConnection().(gorm.DB)

	rows, err := DBConnection.Find(&users).Rows()

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	if err != nil {
		return nil
	}

	var user entities.UserEntity

	for rows.Next() {
		err := DBConnection.ScanRows(rows, &user)
		if err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users

}

func (userRepo *GormUserRepo) DeleteUser(id int) (bool, error) {
	var user entities.UserEntity

	DBConnection := userRepo.DBManager.ProvideDBConnection().(gorm.DB)

	result := DBConnection.Delete(&user, id)

	if result.RowsAffected != 1 {
		errorMessage := fmt.Sprintf("此ID: %d不存在", id)
		return false, errors.New(errorMessage)
	}

	return true, nil
}

func (userRepo *GormUserRepo) GetUser(id int) *entities.UserEntity {
	var user entities.UserEntity
	DBConnection := userRepo.DBManager.ProvideDBConnection().(gorm.DB)
	result := DBConnection.First(&user, id)

	if result.RowsAffected != 1 {
		return nil
	}

	err := result.Row().Scan(
		&user.Id, &user.Vendor, &user.Account, &user.AccountType,
		&user.Password, &user.Name, &user.Status, &user.CreateTime, &user.UpdateTime)

	if err != nil {
		return nil
	}

	return &user
}

func (userRepo *GormUserRepo) UpdateUser(id int, name string) (bool, error) {
	DBConnection := userRepo.DBManager.ProvideDBConnection().(gorm.DB)
	result := DBConnection.Model(&entities.UserEntity{}).Where("id = ?", id).Update("name", name)
	fmt.Println(result.RowsAffected)
	if result.RowsAffected != 1 {
		errorMessage := fmt.Sprintf("此ID: %d不存在或Name: %s，沒有任何變更。", id, name)
		return false, errors.New(errorMessage)
	}

	return true, nil
}

func NewGormUserRepo(dbManager dbManager.InterfaceDBManger) *GormUserRepo {
	return &GormUserRepo{DBManager: dbManager}
}
