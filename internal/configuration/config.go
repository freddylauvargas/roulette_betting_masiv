package configuration

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"sync"
	"log"
)

type config struct {
	DBConnection string
	DBHost       string
	DBInstance   string
	DBPort       string
	DBDatabase   string
	DBUserName   string
	DBPassword   string

	AppServiceName    string
	AppPort           string
	AppPatchLog       string
	LogReviewInterval string

	UrlBetting  	       string
}
func (c *config) GetLogReviewInterval() int {
	i, _ := strconv.Atoi(c.LogReviewInterval)
	return i
}
var (
	// c variable de tipo config que será devuelta
	c    config
	once sync.Once
)
func FromFile() config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading ..env file")
			panic(err)
		}
		c.DBConnection = os.Getenv("DB_CONNECTION")
		c.DBHost = os.Getenv("DB_HOST")
		c.DBInstance = os.Getenv("DB_INSTANCE")
		c.DBPort = os.Getenv("DB_PORT")
		c.DBDatabase = os.Getenv("DB_DATABASE")
		c.DBUserName = os.Getenv("DB_USERNAME")
		c.DBPassword = os.Getenv("DB_PASSWORD")

		c.AppServiceName = os.Getenv("APP_SERVICE_NAME")
		c.AppPort = os.Getenv("APP_PORT")
		c.AppPatchLog = os.Getenv("APP_PATH_LOG")
		c.LogReviewInterval = os.Getenv("LOG_REVIEW_INTERVAL")

		c.UrlBetting = os.Getenv("URL_SIFI_ACTUALIZAR_CLIENTE")
	})
	return c
}