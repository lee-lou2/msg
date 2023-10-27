package sender

import (
	"fmt"
	"github.com/lee-lou2/msg/configs"
	"github.com/lee-lou2/msg/models"
	"regexp"
)

// validate 메시지 유효성 검사
func validate(m *models.Message) ([][]*models.User, error) {
	var maxUser int
	var validateFunc func(u *models.User, m *models.Message) error
	groups := make([][]*models.User, 0)

	// 전체 검증
	if m.Content == "" {
		return groups, fmt.Errorf("내용이 존재하지 않습니다.")
	}
	// 수신자 정보 조회
	if len(m.Users) == 0 {
		return groups, fmt.Errorf("사용자가 존재하지 않습니다.")
	}

	// 개별 검증
	switch m.MessageType {
	case models.SMS.Value():
		validateFunc = validateSMS
		maxUser = configs.SmsMaxUser
	case models.Email.Value():
		validateFunc = validateEmail
		maxUser = configs.EmailMaxUser
	case models.Push.Value():
		validateFunc = validatePush
		maxUser = configs.PushMaxUser
	}
	for _, u := range m.Users {
		// 메세지 전송
		if err := validateFunc(u, m); err != nil {
			// TODO Hook : 개별 유효성 검사 실패시 훅을 남기고 다른 데이터 검증 시작
			continue
		}
		// 묶어서 그룹에 넣기
		if len(groups) == 0 || len(groups[len(groups)-1]) == maxUser {
			groups = append(groups, make([]*models.User, 0))
		}
		groups[len(groups)-1] = append(groups[len(groups)-1], u)
	}
	return groups, nil
}

// validateSMS 메시지 유효성 검사
func validateSMS(u *models.User, m *models.Message) error {
	// 폰번호 확인
	if u.Phone == "" {
		return fmt.Errorf("폰번호가 존재하지 않습니다.")
	}
	// 정규식을 이용한 폰번호 검증
	u.Phone = regexp.MustCompile(`[^0-9]`).ReplaceAllString(u.Phone, "")
	// 전화번호 형식 확인
	regex := regexp.MustCompile(`^01[0-9]{8,9}$`)
	if !regex.MatchString(u.Phone) {
		return fmt.Errorf("전화번호 형식이 올바르지 않습니다")
	}
	return nil
}

// validateEmail 메시지 유효성 검사
func validateEmail(u *models.User, m *models.Message) error {
	// 이메일 존재 여부 확인
	if u.Email == "" {
		return fmt.Errorf("이메일이 존재하지 않습니다.")
	}
	// 이메일 주소 확인
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !regex.MatchString(u.Email) {
		return fmt.Errorf("이메일 형식이 올바르지 않습니다")
	}
	// 제목이 포함되어있는지 확인
	if m.Title == "" {
		return fmt.Errorf("제목이 존재하지 않습니다.")
	}
	return nil
}

// validatePush 메시지 유효성 검사
func validatePush(u *models.User, m *models.Message) error {
	// 메세지 전송
	if u.FcmToken == "" {
		return fmt.Errorf("FCM 토큰이 존재하지 않습니다.")
	}
	// FCM 토큰 형식 확인
	regex := regexp.MustCompile(`^[a-zA-Z0-9\-_]+:[a-zA-Z0-9\-_]+$`)
	if !regex.MatchString(u.FcmToken) {
		return fmt.Errorf("FCM 토큰 형식이 올바르지 않습니다")
	}
	return nil
}
