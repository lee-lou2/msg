package configs

const (
	// 초당 요청 횟수
	EmailRPS = 10000
	SmsRPS   = 10000
	PushRPS  = 10000
	// 한 번에 그룹핑해서 전송할 수 있는 수
	EmailMaxUser = 1
	SmsMaxUser   = 1
	PushMaxUser  = 10
)
