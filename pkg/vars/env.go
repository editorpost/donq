package vars

import (
	"github.com/caarlos0/env/v10"
	"log/slog"
	"sync"
)

var (
	_env     = &Windmill{}
	_envOnce = &sync.Once{}
)

// FromEnv loads the Windmill config instance from the environment.
// Variables loaded with sync.Once.
func FromEnv() *Windmill {

	_envOnce.Do(func() {
		if err := env.Parse(_env); err != nil {
			panic(err)
		}
	})

	return _env
}

// LoggerAttr returns a slice of slog.Attr with Windmill attributes
func LoggerAttr(traceID string) []slog.Attr {

	wm := FromEnv()

	return []slog.Attr{
		slog.String("trace_id", traceID),
		slog.String("wm_job_id", wm.GetRootFlowJobID()),
		slog.String("wm_flow_path", wm.GetFlowPath()),
		slog.String("wm_flow_job_id", wm.GetFlowJobID()),
		slog.String("wm_job_path", wm.GetJobPath()),
	}
}
