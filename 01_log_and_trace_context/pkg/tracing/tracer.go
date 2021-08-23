package tracing

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

// SetupGlobalJaegerTracer sets global tracer client using Jaeger ENVs.
// Refer to https://github.com/jaegertracing/jaeger-client-go for details.
func SetupGlobalJaegerTracer() (io.Closer, error) {
	cfg, err := jaegerCfg.FromEnv()
	if err != nil {
		return nil, fmt.Errorf("jaegerCfg.FromEnv: %w", err)
	}

	if cfg.ServiceName == "" {
		cfg.ServiceName = "Practicum example 1"
	}
	cfg.Sampler.Type = jaeger.SamplerTypeConst
	cfg.Sampler.Param = 1

	tracer, closer, err := cfg.NewTracer(jaegerCfg.Logger(jaegerLog.NullLogger), jaegerCfg.Metrics(metrics.NullFactory))
	if err != nil {
		return nil, fmt.Errorf("creating Jaeger tracer: %w", err)
	}

	opentracing.SetGlobalTracer(tracer)

	return closer, nil
}
