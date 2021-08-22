package tracer

import (
	"github.com/opentracing/opentracing-go"
	cfg "github.com/ozoncp/ocp-certificate-api/internal/config"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

// InitTracer - init tracer
func InitTracer(serviceName string) io.Closer {
	cfgMetrics := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.GetConfigInstance().Jaeger.Host + cfg.GetConfigInstance().Jaeger.Port,
		},
	}
	tracer, closer, err := cfgMetrics.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Err(err).Msgf("failed init jaeger: %v", err)
	}
	opentracing.SetGlobalTracer(tracer)
	log.Info().Msgf("Traces started")

	return closer
}
