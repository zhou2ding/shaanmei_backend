package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"shaanmei_backend/pkg/conf"
	"shaanmei_backend/pkg/logger"
	"time"
)

var dbInstance *gorm.DB

func InitIotDB() error {
	logger.Logger.Infof("connect to database %s", conf.Conf.GetString("database.freight_index.link"))
	db, err := gorm.Open(mysql.Open(conf.Conf.GetString("database.freight_index.link")), &gorm.Config{})
	if err != nil {
		logger.Logger.Fatalf("init db err %s", err)
		return err
	}
	if conf.Conf.GetBool("log.stdout") {
		db.Logger = gormLogger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		})
	}
	dbInstance = db
	return nil
}

func GetDB() *gorm.DB {
	return dbInstance
}
