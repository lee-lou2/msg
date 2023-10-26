package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"os"
)

// Run REST API 실행
func Run() {
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
	})

	// Static
	app.Static("/", "./web")
	// Middleware
	app.Use(pprof.New())
	app.Get("/metrics", monitor.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(recover.New())
	// CORS 설정
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://lou2.kr, http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routers(app)

	log.Println("🚀 [REST API] 프로그램이 시작되었습니다.")
	log.Fatal(app.Listen(":" + os.Getenv("SERVER_PORT")))
}
