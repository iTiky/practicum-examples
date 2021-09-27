package main

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

type Service struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

func (s *Service) Inc(value int) int {
	s.logger.Info("Inc", zap.Int("value", value))
	return value + 1
}

func TestSimplePanic(t *testing.T) {
	s := New(nil)
	res := s.Inc(1)
	assert.Equal(t, 2, res)
}

func TestSimpleOk(t *testing.T) {
	s := New(zap.NewNop())
	res := s.Inc(1)
	assert.Equal(t, 2, res)
}
