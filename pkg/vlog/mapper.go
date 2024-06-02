package vlog

import (
	"log/slog"
)

func Mapper(r slog.Record) map[string]any {

	m := map[string]any{
		"_time": r.Time.UTC(),
		"_msg":  r.Message,
		"level": r.Level.String(),
	}

	r.Attrs(func(a slog.Attr) bool {
		m[a.Key] = a.Value.Any()
		return true
	})

	return m
}

type JobMap interface {
	GetRootFlowJobID() string
	GetFlowPath() string
	GetFlowJobID() string
	GetJobPath() string
}

func WindmillAttr(traceID string, wm JobMap) []slog.Attr {
	return []slog.Attr{
		slog.String("_stream", "trace_id"),
		slog.String("trace_id", traceID),
		slog.String("wm_job_id", wm.GetRootFlowJobID()),
		slog.String("wm_flow_path", wm.GetFlowPath()),
		slog.String("wm_flow_job_id", wm.GetFlowJobID()),
		slog.String("wm_job_path", wm.GetJobPath()),
	}
}
