package dispatcher

import (
	"github.com/lee-lou2/msg/cmd/collector"
	"github.com/lee-lou2/msg/models"
	"github.com/lee-lou2/msg/pkg/orm"
	"log"
	"time"
)

// Run 메시지 생성
func Run() {
	log.Println("🚀 [Dispatcher] 프로그램이 시작되었습니다.")
	for {
		<-time.After(time.Second * 10)
		db, err := orm.GetDB()
		if err != nil {
			panic(err)
		}
		// 현재 시간을 Asia/Seoul 시간대로 설정
		loc, err := time.LoadLocation("Asia/Seoul")
		if err != nil {
			panic(err)
		}
		now := time.Now().In(loc)
		// SendAt이 현재 시간보다 이전이거나 같은 레코드 찾기
		var collections []models.Collection
		err = db.Where("send_at <= ?", now).Find(&collections).Error
		if err != nil {
			panic(err)
		}

		// 찾은 레코드의 PK 값들을 수집
		var ids []uint
		for _, collection := range collections {
			ids = append(ids, collection.ID)
		}

		if len(ids) > 0 {
			// 찾은 레코드의 PK 값들을 사용하여 레코드 삭제
			err = db.Delete(&models.Collection{}, ids).Error
			if err != nil {
				panic(err)
			}
		}

		// 전송
		go func() {
			for _, collection := range collections {
				collector.RawMsg <- &collection
			}
		}()
	}
}
