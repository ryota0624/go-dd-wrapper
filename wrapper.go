package ddwrapper

import (
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func RunWithSpan[T any](ctx context.Context, spanName string, body func(context.Context) (T, error)) (T, error) {
	span, ctx_ := tracer.StartSpanFromContext(ctx, spanName)
	result, err := body(ctx_)
	span.Finish(func(cfg *ddtrace.FinishConfig) {
		if err != nil {
			cfg.Error = err
		}
	})
	return result, err
}

func RunWithSpanNoError[T any](ctx context.Context, spanName string, body func(context.Context) T) T {
	ret, err := RunWithSpan(ctx, spanName, func(ctx context.Context) (T, error) {
		ret := body(ctx)
		return ret, nil
	})

	if err != nil {
		panic(err)
	}

	return ret
}
