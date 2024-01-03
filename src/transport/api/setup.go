package api

import (
	"fmt"
	"guide_go/src/application/service"
	"guide_go/src/domain"
	"guide_go/src/internal/launch"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type APIMethod interface {
	GET(path string, f func(ApiContext) error, name ...string)
	DELETE(path string, f func(ApiContext) error, name ...string)
	PUT(path string, f func(ApiContext) error, name ...string)
	POST(path string, f func(ApiContext) error, name ...string)
}

type ApiContext struct {
	Fiber *fiber.Ctx
	Name  string
	Path  string
	Env   *domain.Environment
}

type ApiLauncher struct {
	Fiber *fiber.App
	*launch.Launcher
	ServicePool *service.ServicePool
}

type APIStandardLauncher struct {
	ApiLauncher
}

func ServeHttp(launcher *launch.Launcher) {
	defer func() {
		if r := recover(); r != nil {

			os.Exit(0)
		}
	}()
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Failed to create rally-gateway logger: %v", err)
	}
	defer logger.Sync()
	al := ApiLauncher{}

	al.Launcher = launcher
	al.ServicePool = service.NewServicePool(launcher)
	al.Fiber = fiber.New(fiber.Config{
		AppName:   "Guide",
		BodyLimit: 50 * 1024 * 1024,
	})
	al.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: func() string {

			return "*"

		}(),
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodDelete,
			fiber.MethodPut}, ","),
	}))
	al.Fiber.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		// Request 처리
		err := c.Next()
		// 로그 기록
		logger.Info("guide-logs",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", time.Since(start)),
		)
		return err
	})
	APIStandardLauncher{al}.APISetUp()
	// fmt.Println("launcher.Env.ApiBase.Port", launcher.Env.ApiBase.Port)
	al.Fiber.Listen(fmt.Sprintf("0.0.0.0:%d", launcher.Env.ApiBase.Port))
}
