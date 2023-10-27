package firebase

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/lee-lou2/msg/models"
	"google.golang.org/api/option"
	"io/ioutil"
	"os"
	"strings"
)

var client *messaging.Client
var ctx = context.Background()

// SendPush 푸시 전송
func SendPush(users []*models.User, title, body string) (int, int, error) {
	tokens := make([]string, len(users))
	for i, u := range users {
		tokens[i] = u.FcmToken
	}
	if client == nil {
		serviceAccountKeyFilePath := "serviceAccountKey.json"
		serviceAccountKeyValue := os.Getenv("FCM_SERVICE_ACCOUNT_KEY")

		// 파일이 이미 존재하는지 확인
		if _, err := os.Stat(serviceAccountKeyFilePath); os.IsNotExist(err) {
			// 파일이 존재하지 않으면 JSON 문자열을 파일에 저장
			var jsonValue map[string]interface{}
			err := json.NewDecoder(strings.NewReader(serviceAccountKeyValue)).Decode(&jsonValue)
			if err != nil {
				panic(err) // 적절한 오류 처리를 추가하세요.
			}

			jsonData, err := json.MarshalIndent(jsonValue, "", "  ")
			if err != nil {
				panic(err) // 적절한 오류 처리를 추가하세요.
			}

			err = ioutil.WriteFile(serviceAccountKeyFilePath, jsonData, 0644)
			if err != nil {
				panic(err) // 적절한 오류 처리를 추가하세요.
			}
		}
		opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			return 0, 0, err
		}
		client, err = app.Messaging(ctx)
		if err != nil {
			return 0, 0, err
		}
	}
	// 메시지 지정
	messages := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Tokens: tokens,
	}
	resp, err := client.SendMulticast(ctx, messages)
	if err != nil {
		if resp == nil {
			return 0, 0, err
		}
		return resp.SuccessCount, resp.FailureCount, err
	}
	if resp.FailureCount > 0 {
		return resp.SuccessCount, resp.FailureCount, fmt.Errorf("failed to send message")
	}
	return resp.SuccessCount, resp.FailureCount, nil
}
