package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

	address := os.Getenv("DUMMY_ADDRESS")
	produceLogDuration := os.Getenv("DUMMY_PRODUCE_LOG_DURATION")

	if address == "" {
		log.Fatal("No address!")
	}

	if produceLogDuration == "" {
		log.Fatal("No Produce Every!")
	}

	logDuration, err := time.ParseDuration(produceLogDuration)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Dummy Server start!")

	// production log
	produceLog(logDuration)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		log.Info("Dummy got request")
		return c.String(http.StatusOK, "Yes, I am dummy")
	})

	// start the service
	err = e.Start(":1323")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func produceLog(duration time.Duration) {
	ticker := time.NewTicker(duration)
	go func() {
		for t := range ticker.C {
			log.Errorf("Dummy error at %s", t)
		}
	}()
}
