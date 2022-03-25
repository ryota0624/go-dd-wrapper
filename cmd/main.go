package main

import (
	"context"
	"log"

	ddwrapper "github.com/ryota0624/dd-wrapper"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func main() {
	err := profiler.Start(
		profiler.WithService("<SERVICE_NAME>"),
		profiler.WithEnv("<ENVIRONMENT>"),
		profiler.WithVersion("<APPLICATION_VERSION>"),
		profiler.WithTags("<KEY1>:<VALUE1>,<KEY2>:<VALUE2>"),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
			// The profiles below are disabled by default to keep overhead
			// low, but can be enabled as needed.

			// profiler.BlockProfile,
			// profiler.MutexProfile,
			// profiler.GoroutineProfile,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

	tracer.Start()
	defer tracer.Stop()

	_, ctx := tracer.StartSpanFromContext(context.Background(), "root")
	culc := ddwrapper.RunWithSpanNoError(ctx, "culc", func(ctx context.Context) int {
		return 10
	})

	println(culc)
}
