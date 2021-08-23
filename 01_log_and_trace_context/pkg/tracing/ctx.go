package tracing

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// StartSpanFromCtx starts a new Span and returns an enriched Context.
// CorrelationID option is added by default.
func StartSpanFromCtx(ctx context.Context, opName string, opts ...SpanOption) (context.Context, opentracing.Span) {
	if ctx == nil {
		ctx = context.Background()
	}

	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		childSpan := opentracing.StartSpan(opName, opentracing.ChildOf(parentSpan.Context()))
		return newCtxWithSpanOptions(ctx, childSpan, opts...)
	}

	newSpan := opentracing.StartSpan(opName)
	opts = append(opts, CorrelationIDOpt(ctx))

	return newCtxWithSpanOptions(ctx, newSpan, opts...)
}

// FinishSpan ends Span, sets Span tags and logs for operation result.
func FinishSpan(span opentracing.Span, err error) {
	if span == nil {
		return
	}
	defer span.Finish()

	if err != nil {
		ext.LogError(span, err)
	} else {
		ext.Error.Set(span, false)
	}
}

// newCtxWithSpanOptions updates Span with SpanOption options and returns enriched context.
func newCtxWithSpanOptions(ctx context.Context, span opentracing.Span, opts ...SpanOption) (context.Context, opentracing.Span) {
	for _, opt := range opts {
		opt(span)
	}

	return opentracing.ContextWithSpan(ctx, span), span
}
