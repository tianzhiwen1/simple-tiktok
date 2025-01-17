package db

import (
	"simple-tiktok/pkg/consts"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond * 100,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)
	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt: true,
			Logger:      gormlogrus,
		},
	)
	if err != nil {
		panic(err)
	}
	if err := DB.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}

	if !DB.Migrator().HasTable(&User{}) {
		if err := DB.AutoMigrate(&User{}); err != nil {
			panic(err)
		}
	}

	if !DB.Migrator().HasTable(&Video{}) {
		if err := DB.AutoMigrate(&Video{}); err != nil {
			panic(err)
		}
	}

	if !DB.Migrator().HasTable(&Comment{}) {
		if err := DB.AutoMigrate(&Comment{}); err != nil {
			panic(err)
		}
	}
}
