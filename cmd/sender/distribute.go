package sender

import (
	"fmt"
	"github.com/lee-lou2/msg/models"
	"github.com/lee-lou2/msg/pkg/email"
	"github.com/lee-lou2/msg/pkg/firebase"
	"github.com/lee-lou2/msg/pkg/twilio"
	"sync/atomic"
	"time"
)

// distribute 메세지 분배
func distribute(m *models.Message, tickers map[int]*time.Ticker) error {
	msgType := m.MessageType
	groups, err := validate(m)
	if err != nil {
		return err
	}
	switch msgType {
	case models.SMS.Value():
		<-tickers[models.SMS.Value()].C
		if m.IsTest {
			atomic.AddInt32(&Counter, 1)
		} else {
			go func() {
				// 리스트안에 리스트가 있는 변수
				for _, g := range groups {
					if err := twilio.SendSMSTwilio(g, m.Content); err != nil {
						// TODO Hook
					}
				}
			}()
		}
	case models.Email.Value():
		<-tickers[models.Email.Value()].C
		if m.IsTest {
			atomic.AddInt32(&Counter, 1)
		} else {
			go func() {
				// 리스트안에 리스트가 있는 변수
				for _, g := range groups {
					if err := email.SendEmail(g, m.Title, m.Content); err != nil {
						// TODO Hook
					}
				}
			}()
		}
	case models.Push.Value():
		<-tickers[models.Push.Value()].C
		if m.IsTest {
			atomic.AddInt32(&Counter, 1)
		} else {
			go func() {
				// 리스트안에 리스트가 있는 변수
				for _, g := range groups {
					success, fail, err := firebase.SendPush(g, m.Title, m.Content)
					if err != nil {
						// TODO Hook
						fmt.Println(success, fail, err)
					}
				}
			}()
		}
	}
	return nil
}
