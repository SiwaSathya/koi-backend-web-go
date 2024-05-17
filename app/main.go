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
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	detKegRepo := repository.NewPostgreDetailKegiatan(db.GormClient.DB)
	evntRepo := repository.NewPostgreEvent(db.GormClient.DB)
	metKegRepo := repository.NewPostgreMetodePembayaran(db.GormClient.DB)
	naraRepo := repository.NewPostgreNarahubung(db.GormClient.DB)
	pmbyrnRepo := repository.NewPostgrePembayaran(db.GormClient.DB)

	timeoutContext := fiber.Config{}.ReadTimeout

	userUseCase := usecase.NewUserUseCase(usrRepo, ormRepo, mhsRepo, timeoutContext)
	eventUseCase := usecase.NewEventUseCase(detKegRepo, evntRepo, naraRepo, metKegRepo, usrRepo, ormRepo, timeoutContext)
	detKegUseCase := usecase.NewDetailKegiatanUseCase(detKegRepo, timeoutContext)
	pmbyrnUseCase := usecase.NewPembayaranUseCase(pmbyrnRepo, mhsRepo, timeoutContext)

	app := fiber.New(fiber.Config{})
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${green} ${status} ${white} | ${latency} | ${ip} | ${green} ${method} ${white} | ${path} | ${yellow} ${body} ${reset} | ${magenta} ${resBody} ${reset}\n",
		TimeFormat: "02 January 2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))
	app.Use(

		func(c *fiber.Ctx) error {
			return c.Next()
		},

		cors.New(cors.Config{
			AllowCredentials: true,
			AllowOrigins:     "http://localhost:*.*",
			AllowHeaders:     "Accept, Authorization, Content-Type, Origin, Referer, User-Agent, X-Requested-With, Accept-Encoding, Accept-Language",
			AllowMethods:     "*",
			MaxAge:           0,
			ExposeHeaders:    "Content-Disposition",
		}),
	)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {

		//call delivery http here
		delivery.NewHealthCheckHandler(app)
		delivery.NewUserHandler(app, userUseCase)
		delivery.NewEventHandler(app, eventUseCase)
		delivery.NewDetailKegiatanHandler(app, detKegUseCase)
		delivery.NewPembayaranHandler(app, pmbyrnUseCase)
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
