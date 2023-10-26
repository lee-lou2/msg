package collector

import (
	"fmt"
	"github.com/lee-lou2/msg/models"
	"log"
)

var data chan *models.Collection
var RawMsg chan *models.Collection

// Run 메시지 수집기 실행
func Run() {
	RawMsg = make(chan *models.Collection)
	log.Println("🚀 [Collector] 프로그램이 시작되었습니다.")
	// 하나씩 메세지 저장
	go func() {
		for {
			c := <-data
			// 메세지 전송
			if err := c.Create(); err != nil {
				log.Println(fmt.Errorf("메세지 저장간 오류 발생, 오류 내용 : %w", err))
			}
		}
	}()
	// 메세지 수집
	for {
		c := <-RawMsg
		if err := distribute(c); err != nil {
			log.Println(fmt.Errorf("메세지 수집간 오류 발생, 오류 내용 : %w", err))
		}
	}
}
