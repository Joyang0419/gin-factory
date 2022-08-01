package userRepo

import (
	"aCupOfGin/internal/entities"
	"aCupOfGin/internal/tools/dbManager"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type CSVUserRepo struct {
	DBManager dbManager.InterfaceDBManger
}

func (userRepo *CSVUserRepo) CreateUser(
	vendor string,
	account string,
	accountType string,
	hashedPassword string,
	name string) (bool, error) {

	fileName := userRepo.DBManager.ProvideDBConnection().(string)

	file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0777)

	reader := csv.NewReader(file)

	records, _ := reader.ReadAll()

	user := entities.UserEntity{
		Id:          len(records) + 1,
		Vendor:      vendor,
		Account:     account,
		AccountType: accountType,
		Password:    hashedPassword,
		Name:        name,
		Status:      "active",
	}

	currentTime := time.Now().Format(time.RFC3339) //time.RFC3339 is the format about ISO8601

	writer := csv.NewWriter(file)

	rowData := []string{strconv.Itoa(user.Id), user.Vendor, user.Account, user.AccountType, user.Password, user.Name, user.Status, currentTime, currentTime}

	err := writer.Write(rowData)

	if err != nil {
		return false, err
	}

	defer func() {
		writer.Flush()
		if err := writer.Error(); err != nil {
			fmt.Println(err)
		}
	}()

	return true, nil
}

func (userRepo *CSVUserRepo) GetAllUsers() []entities.UserEntity {
	var users []entities.UserEntity

	fileName := userRepo.DBManager.ProvideDBConnection().(string)

	file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0777)

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		userId, _ := strconv.Atoi(record[0])
		userCreateTime, _ := time.Parse(time.RFC3339, record[7])
		userUpdateTime, _ := time.Parse(time.RFC3339, record[8])

		user := entities.UserEntity{
			Id:          userId,
			Vendor:      record[1],
			Account:     record[2],
			AccountType: record[3],
			Password:    record[4],
			Name:        record[5],
			Status:      record[6],
			CreateTime:  userCreateTime,
			UpdateTime:  userUpdateTime,
		}
		users = append(users, user)

	}

	return users

}

func (userRepo *CSVUserRepo) GetUser(id int) *entities.UserEntity {
	var user entities.UserEntity
	fileName := userRepo.DBManager.ProvideDBConnection().(string)

	file, _ := os.OpenFile(fileName, os.O_RDWR, 0777)

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		userId, _ := strconv.Atoi(record[0])

		if userId == id {
			userCreateTime, _ := time.Parse(time.RFC3339, record[7])
			userUpdateTime, _ := time.Parse(time.RFC3339, record[8])

			user = entities.UserEntity{
				Id:          userId,
				Vendor:      record[1],
				Account:     record[2],
				AccountType: record[3],
				Password:    record[4],
				Name:        record[5],
				Status:      record[6],
				CreateTime:  userCreateTime,
				UpdateTime:  userUpdateTime,
			}

			return &user
		}

	}

	return nil
}

func (userRepo *CSVUserRepo) DeleteUser(id int) (bool, error) {
	fileName := userRepo.DBManager.ProvideDBConnection().(string)

	file, _ := os.OpenFile(fileName, os.O_RDWR, 0777)

	reader := csv.NewReader(file)

	var records [][]string

	deleteStatus := false

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		userId, _ := strconv.Atoi(record[0])

		if userId == id {
			deleteStatus = true
			continue
		}

		records = append(records, record)

	}

	fileErr := file.Close()
	if fileErr != nil {
		return false, fileErr
	}

	newFile, _ := os.Create(fileName)
	writer := csv.NewWriter(newFile)

	err := writer.WriteAll(records)

	if err != nil {
		return false, err
	}

	if !deleteStatus {
		return false, nil
	}

	return true, nil
}

func (userRepo *CSVUserRepo) UpdateUser(id int, name string) (bool, error) {
	fileName := userRepo.DBManager.ProvideDBConnection().(string)

	file, _ := os.OpenFile(fileName, os.O_RDWR, 0777)

	reader := csv.NewReader(file)

	var records [][]string

	updateStatus := false

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		userId, _ := strconv.Atoi(record[0])

		if userId == id {
			updateStatus = true
			record[5] = name
		}

		records = append(records, record)

	}

	fileErr := file.Close()
	if fileErr != nil {
		return false, fileErr
	}

	newFile, _ := os.Create(fileName)
	writer := csv.NewWriter(newFile)

	err := writer.WriteAll(records)

	if err != nil {
		return false, err
	}

	if !updateStatus {
		return false, nil
	}

	return true, nil
}

func NewCSVUserRepo(dbManager dbManager.InterfaceDBManger) *CSVUserRepo {
	return &CSVUserRepo{DBManager: dbManager}
}
