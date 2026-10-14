package database

import (
	"crud/internal/config"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Config) *gorm.DB {
	logLevel := logger.Warn
	if cfg.AppEnv == "development" {
		logLevel = logger.Info
	}

	// Mở kết nối
	db, err := gorm.Open(postgres.Open(cfg.DB.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		log.Fatal("Không kết nối được database: %v", err)
	}

	// Lấy *sql.DB bên dưới GORM để chỉnh connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("lỗi lấy sql.DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)           // số kết nối rảnh được giữ lại để tái sử dụng
	sqlDB.SetMaxOpenConns(50)           // số kết nối tối đa mở cùng lúc
	sqlDB.SetConnMaxLifetime(time.Hour) // sau 1 giờ thì đóng và tạo kết nối mới

	return db
}
