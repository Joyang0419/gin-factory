package dbManager

import (
	"log"
	"os"
)

type FileName string

type CSVDBMSetting struct {
	FileName FileName
}

type CSVDBManager struct {
	Settings *CSVDBMSetting
}

func (Manager *CSVDBManager) Init() {
	file, err := os.OpenFile(string(Manager.Settings.FileName), os.O_CREATE, 0777)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
}

func (Manager *CSVDBManager) IsConnected() bool {
	_, err := os.OpenFile(string(Manager.Settings.FileName), os.O_RDONLY, 0777)
	if err != nil {
		return false
	}
	return true
}

func (Manager *CSVDBManager) ProvideDBConnection() any {
	return string(Manager.Settings.FileName)
}

func NewCSVDBMSetting(fileName FileName) *CSVDBMSetting {
	return &CSVDBMSetting{
		FileName: fileName,
	}
}

func NewCSVDBManager(dbmSetting *CSVDBMSetting) *CSVDBManager {
	manager := CSVDBManager{Settings: dbmSetting}
	manager.Init()

	return &manager
}
