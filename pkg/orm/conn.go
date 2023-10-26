package orm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDB DB 연결
func GetDB() (*gorm.DB, error) {
	var err error
	if db == nil {
		db, err = gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		// 커넥션 풀 적용
		sqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}
		sqlDB.SetMaxIdleConns(100)
		sqlDB.SetMaxOpenConns(1000)
	}
	return db, nil
}
