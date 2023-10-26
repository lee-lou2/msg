package models

import (
	"github.com/lee-lou2/msg/pkg/orm"
)

func init() {
	db, err := orm.GetDB()
	if err != nil {
		panic(err)
	}
	// 테이블 생성
	if err := db.AutoMigrate(&Collection{}); err != nil {
		panic(err)
	}
}
