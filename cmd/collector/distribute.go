package collector

import (
	"encoding/json"
	"fmt"
	"github.com/lee-lou2/msg/cmd/sender"
	"github.com/lee-lou2/msg/models"
	"time"
)

// distribute 메시지 분배
func distribute(c *models.Collection) error {
	m := &models.Message{}
	if err := json.Unmarshal([]byte(c.Raw), m); err != nil {
		return err
	}

	// 타입 지정
	m.MessageType = c.MsgType

	// 발송 시간 조회
	sendAtStr := c.SendAt
	if sendAtStr == "" {
		sendAtStr = m.SendAt
	}
	if sendAtStr == "" {
		return fmt.Errorf("invalid send_at: %s", sendAtStr)
	}
	// 발송 시간 검증
	sendAt, err := time.Parse("2006-01-02 15:04:05", sendAtStr)
	if err != nil {
		return err
	}
	// UTC 로 변환 후 검증
	sendAt = sendAt.Add(-9 * time.Hour)

	// 반복 횟수 확인
	if m.IntervalCount == 0 {
		m.IntervalCount = 1
	}
	for i := 0; i < m.IntervalCount; i++ {
		// 발송 시간 계산
		if i != 0 {
			sendAt = sendAt.Add(time.Duration(m.Interval) * time.Second)
		}
		if sendAt.Before(time.Now().UTC()) {
			// 즉시 발송
			sender.Msg <- m
		} else {
			// 예약 발송
			sendAt = sendAt.Add(9 * time.Hour)
			sendAtStr := sendAt.Format("2006-01-02 15:04:05")
			raw, err := json.Marshal(&models.Message{
				Users:   m.Users,
				Title:   m.Title,
				Content: m.Content,
				SendAt:  sendAtStr,
			})
			if err != nil {
				return err
			}
			data <- &models.Collection{
				MsgType: c.MsgType,
				Raw:     string(raw),
				SendAt:  sendAtStr,
			}
		}
	}
	return nil
}
