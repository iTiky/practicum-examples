package main

import (
	"github.com/itiky/practicum-examples/02_zap_logger/logger"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

type Event struct {
	Id   int
	Type string
	Time time.Time
}

type JSONEvent struct {
	Id   int       `json:"id"`
	Type string    `json:"type"`
	Time time.Time `json:"time"`
}

func main() {
	// создаем логгер
	logger, err := logger.New("info")
	if err != nil {
		log.Fatal(err)
	}

	event := Event{1, "CreateEvent", time.Now()}

	// выводим по полям
	logger.Info("created eventservice",
		zap.Int("EventId", event.Id),
		zap.String("EventType", event.Type),
		zap.Time("EventTime", event.Time),
	)

	logger.Info("created eventservice",
		zap.Reflect("Event", event),
	)

	logger.Debug("check debug")

	logger.Info("check info")

	logger.Warn("check warn")

	logger.Error("check error")

	jsonEvent := JSONEvent{1, "CreateEvent", time.Now()}

	logger = logger.With(zap.String("traceId", "uniq_trace_id"))

	logger.Info("created eventservice",
		zap.Reflect("Event", jsonEvent),
	)

	// Завершает приложение
	_, err = os.ReadFile("not_exist_file.txt")
	if err != nil {
		logger.Fatal("app exit", zap.Error(err))
	}

}
