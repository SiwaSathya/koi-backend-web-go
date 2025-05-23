package main

import (
	"fmt"
	"koi-backend-web-go/db"
	"koi-backend-web-go/koi/delivery"
	"koi-backend-web-go/koi/repository"
	"koi-backend-web-go/koi/usecase"
	"log"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/pandeptwidyaop/golog"
)

func main() {
	Init()
	// initEnv()
	listenPort := ":4000"
	appName := os.Getenv("APP_NAME")

	mhsRepo := repository.NewPostgreMahasiswa(db.GormClient.DB)
	ormRepo := repository.NewPostgreOrmawa(db.GormClient.DB)
	usrRepo := repository.NewPostgreUser(db.GormClient.DB)
	timeoutContext := fiber.Config{}.ReadTimeout

	userUseCase := usecase.NewLocationUseCase(usrRepo, ormRepo, mhsRepo, timeoutContext)

	app := fiber.New(fiber.Config{})
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${green} ${status} ${white} | ${latency} | ${ip} | ${green} ${method} ${white} | ${path} | ${yellow} ${body} ${reset} | ${magenta} ${resBody} ${reset}\n",
		TimeFormat: "02 January 2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {

		//call delivery http here
		delivery.NewHealthCheckHandler(app)
		delivery.NewUserHandler(app, userUseCase)
		log.Fatal(app.Listen(listenPort))
		wg.Done()
	}()
	golog.Slack.Info(fmt.Sprintf("%s: App Start & Running", appName))
	wg.Wait()
}

func Init() {
	InitEnv()
	InitSlack()
	InitDB()
}

func InitEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}
}

func InitSlack() {
	golog.New()
}

func InitDB() {
	db.NewGormClient()
}
