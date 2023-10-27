package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/lee-lou2/msg/models"
	"google.golang.org/api/option"
	"os"
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
		opt := option.WithCredentialsJSON([]byte(os.Getenv("FCM_SERVICE_ACCOUNT_KEY")))
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
