package db

import (
	"fmt"
	"go-project/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbPg *dbPgsql

type dbPgsql struct {
	dbpg *gorm.DB
}

func GetConnectionDB() (*gorm.DB, error) {
	return DbPg.Getconnection()
}

func (dbpgsql *dbPgsql) Getconnection() (*gorm.DB, error) {
	return dbpgsql.dbpg, nil
}

func InitConnectionDB() error {
	DbPg = new(dbPgsql)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.AppEnv.DBHost, config.AppEnv.DBUserPassword, config.AppEnv.DBUserPassword, config.AppEnv.DBName, config.AppEnv.DBPort)

	conn, errdb := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errdb != nil {
		log.Fatal("Failed to connect database")
	}

	DbPg.dbpg = conn
	return nil
}
