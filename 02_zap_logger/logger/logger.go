package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// New create new logger
func New(level string) (*zap.Logger, error) {
	// создаем конфиг
	cfg := zap.NewProductionConfig()

	// задаем минимальный уровень логирования
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	// задаем уровень логирования приложения
	atom := zap.NewAtomicLevel()
	err := atom.UnmarshalText([]byte(level))
	if err != nil {
		return nil, err
	}
	cfg.Level = atom

	// Куда выводим
	cfg.OutputPaths = []string{"stdout"}

	// Конфигурируем формат выводимого времени (если нужно поменять)
	cfg.EncoderConfig.EncodeTime = CustomMillisTimeEncoder

	return cfg.Build()
}

// CustomMillisTimeEncoder unix ms time encoder
func CustomMillisTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000Z07"))
}
