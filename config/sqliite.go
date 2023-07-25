package config

import (
	"os"

	"github.com/DarlanPrado-Access/primeiro-projeto-go/app/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "dns/path"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Arquivo de base de dados não encontrado, criando...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Erro de conexão com a base: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&model.ModelExemple{})
	if err != nil {
		logger.Errorf("Erro de migração: %v", err)
		return nil, err
	}

	return db, nil
}
