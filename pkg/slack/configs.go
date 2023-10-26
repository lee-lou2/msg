package slack

import (
	"fmt"
	"os"
)

type Configs struct {
	BotToken string
}

// setConfigs 슬랙 설정
func setConfigs(defaultConfigs ...Configs) (*Configs, error) {
	configs := Configs{
		BotToken: os.Getenv("SLACK_BOT_TOKEN"),
	}
	if len(defaultConfigs) > 0 {
		configs = defaultConfigs[0]
	}
	if configs.BotToken == "" {
		return nil, fmt.Errorf("Slack 설정이 잘못되었습니다.")
	}
	return &configs, nil
}
