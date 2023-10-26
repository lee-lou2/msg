package rest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lee-lou2/msg/cmd/collector"
	"github.com/lee-lou2/msg/cmd/sender"
	"github.com/lee-lou2/msg/models"
	"sync/atomic"
)

// sendHandler 메세지 전송
func sendHandler(c *fiber.Ctx) error {
	body := c.Body()
	msgType, err := models.StrToMsgType(c.Params("type"))
	if err != nil {
		return fmt.Errorf("invalid msg_type: %s", c.Params("type"))
	}
	go func() {
		collector.RawMsg <- &models.Collection{
			MsgType: msgType,
			Raw:     string(body),
		}
	}()
	return c.JSON(fiber.Map{"message": "success"})
}

// countHandler 카운터 조회
func countHandler(c *fiber.Ctx) error {
	cnt := atomic.LoadInt32(&sender.Counter)
	return c.JSON(fiber.Map{"count": cnt})
}
