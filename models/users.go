package models

// User 사용자
type User struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	FcmToken string `json:"fcm_token"`
}
