package models

import (
	"github.com/lee-lou2/msg/pkg/orm"
	"gorm.io/gorm"
)

// Collection 수집
type Collection struct {
	gorm.Model
	MsgType int    `json:"msg_type" gorm:"type:int;not null;comment:'메시지 타입'"`
	Raw     string `json:"raw" gorm:"type:text;not null;comment:'원본 데이터'"`
	SendAt  string `json:"send_at" gorm:"type:datetime;not null;index;comment:'발송 시간'"`
}

// Create 수집 생성
func (c *Collection) Create() error {
	db, err := orm.GetDB()
	if err != nil {
		return err
	}
	if err := db.Create(&c).Error; err != nil {
		return err
	}
	return nil
}
