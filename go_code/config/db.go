package config

import (
	"go_code/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to initialize database, got error: %v", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns) // 连接池中最大连接数
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenCons)  // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour)                    // 连接最大生命周期

	if err != nil {
		log.Fatalf("Failed to configure database, got error: %v", err)
	}

	global.Db = db
}
