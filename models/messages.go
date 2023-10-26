package models

import "errors"

// Message 메시지
type Message struct {
	// 사용자 정보
	Users       []*User `json:"users"`
	MessageType int     `json:"message_type"`

	// 메세지 정보
	Title   string `json:"title"`
	Content string `json:"content"`

	// 메세지 발송 정보
	SendAt        string `json:"send_at"`
	Interval      int    `json:"interval"`
	IntervalCount int    `json:"interval_count"`

	// 상태
	IsTest bool `json:"is_test"`
}

type MsgType int

const (
	SMS MsgType = iota + 1
	Push
	Email
)

func (nt MsgType) String() string {
	switch nt {
	case SMS:
		return "sms"
	case Push:
		return "push"
	case Email:
		return "email"
	default:
		return "unknown"
	}
}

func (nt MsgType) Value() int {
	return int(nt)
}

// StrToMsgType 문자열을 메시지 타입으로 변환
func StrToMsgType(s string) (int, error) {
	switch s {
	case "sms":
		return SMS.Value(), nil
	case "push":
		return Push.Value(), nil
	case "email":
		return Email.Value(), nil
	default:
		return 0, errors.New("invalid notification type")
	}
}
