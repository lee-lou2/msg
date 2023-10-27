package sender

import (
	"fmt"
	"github.com/lee-lou2/msg/configs"
	"github.com/lee-lou2/msg/models"
	"log"
	"sync/atomic"
	"time"
)

var count int32
var LastCount int32
var Msg chan *models.Message

// Run 메시지 수집기 실행
func Run() {
	log.Println("🚀 [Sender] 프로그램이 시작되었습니다.")
	Msg = make(chan *models.Message)
	tickers := map[int]*time.Ticker{
		models.SMS.Value():   time.NewTicker(time.Second / time.Duration(configs.SmsRPS)),
		models.Email.Value(): time.NewTicker(time.Second / time.Duration(configs.EmailRPS)),
		models.Push.Value():  time.NewTicker(time.Second / time.Duration(configs.PushRPS)),
	}
	// 초당 카운트 출력
	go func() {
		LastCount = 0
		for {
			<-time.NewTicker(1 * time.Second).C
			cnt := atomic.LoadInt32(&count)
			if cnt != 0 {
				LastCount = cnt
				log.Println("⏱️ RPS : ", cnt)
				atomic.StoreInt32(&count, 0)
			}
		}
	}()
	for {
		m := <-Msg
		if err := distribute(m, tickers); err != nil {
			log.Println(fmt.Errorf("메세지 전송간 오류 발생, 오류 내용 : %w", err))
		}
	}
}
