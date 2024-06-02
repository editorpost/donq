package vlog_test

import (
	"github.com/editorpost/donq/pkg/vlog"
	"log/slog"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	vlog.StdoutLogger(WindmillVars()...)
	slog.Info("test message", slog.String("key", "value"))
	time.Sleep(50 * time.Millisecond)
}

func TestNewElastic(t *testing.T) {
	vlog.VictoriaLogger("", WindmillVars()...)
	slog.Info("test message", slog.String("key", "value"))
	time.Sleep(50 * time.Millisecond)
}

func WindmillVars() []slog.Attr {

	return []slog.Attr{
		slog.String("trace_id", "trace-id"),
		slog.String("job_id", "job-id"),
		slog.String("flow_path", "flow-path"),
		slog.String("flow_job_id", "flow-job-id"),
	}
}
