package mysql

import (
	"log"

	"github.com/Ndraaa15/musiku/cmd/config"
	"github.com/Ndraaa15/musiku/global/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func NewMySqlClient() (*DB, error) {
	db, err := gorm.Open(mysql.Open(config.MySQLConfig()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("[musiku-mysql] failed to connecting with musiku database : %v\n", err)
		return nil, errors.ErrConnectDatabase
	}
	return &DB{db}, nil
}
