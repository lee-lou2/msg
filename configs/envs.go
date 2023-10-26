package configs

import (
	"github.com/joho/godotenv"
	"github.com/lee-lou2/msg/pkg/aws"
)

// LoadEnvs 환경변수 로드
func LoadEnvs() error {
	// .env 파일 로드
	_ = godotenv.Load()
	// 파라미터 스토어를 이용한 환경 변수 설정
	_ = aws.LoadParams()
	return nil
}
