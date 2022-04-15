package main

import (
	// "bytes"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// "context"
	"encoding/json"
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "net/http"
	docs "github.com/DmitriyKalekin/stalker22/docs"
	// "github.com/DmitriyKalekin/stalker22/dto"
	handlers "github.com/DmitriyKalekin/stalker22/handlers"
	telegram "github.com/DmitriyKalekin/stalker22/telegram_api_client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	"time"
	// "strings"
)

const (
	MONGO_DSN        = "mongodb://admin:password@localhost:27017/jirabackup"
	MONGO_DB         = "jirabackup"
	MONGO_COLLECTION = "project"
	TELEGRAM_TOKEN   = "5390968093:AAHm74AS-MopGTOb1LeT6Z_unnkhFeFURA4"
	WH_URL           = "https://2014-45-130-87-17.eu.ngrok.io" + "/tg" + TELEGRAM_TOKEN
)

var (
	mongo_client *mongo.Client = nil
)

func init() {
	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{
		FullTimestamp: true,
	}}
	formatter.Line = true
	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	// mongo_client, err := mongo.Connect(
	// 	context.TODO(),
	// 	options.Client().ApplyURI(MONGO_DSN),
	// )

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// mongo_client = nil
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:3000
// @BasePath  /

func main() {
	log.Debug("Started")

	docs.SwaggerInfo.BasePath = "/"
	port := "3000"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/swagger/doc.json"),
	))

	r.Get("/posts", handlers.ListPosts)
	r.Post("/tg"+TELEGRAM_TOKEN, handlers.TgHandler)

	client := telegram.NewClient(TELEGRAM_TOKEN)

	resp, err := client.SetWebhook(WH_URL)

	if err != nil {
		log.Error(err)
	}
	json_bytes, _ := json.Marshal(resp)

	log.Warnf("%#v", resp)
	log.Debug(string(json_bytes))

	log.Infof("Server started at %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
