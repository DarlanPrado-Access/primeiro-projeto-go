package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgeSQL() (*gorm.DB, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

	// _, err := os.Stat(dbPath)
	// if os.IsNotExist(err) {
	// 	logger.Info("Arquivo de base de dados não encontrado, criando...")
	// 	err = os.MkdirAll("./db", os.ModePerm)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	file, err := os.Create(dbPath)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	file.Close()
	// }

	// db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	// if err != nil {
	// 	logger.Errorf("Erro de conexão com a base: %v", err)
	// 	return nil, err
	// }
	// err = db.AutoMigrate(&model.ModelExemple{})
	// if err != nil {
	// 	logger.Errorf("Erro de migração: %v", err)
	// 	return nil, err
	// }

	// return db, nil
}
