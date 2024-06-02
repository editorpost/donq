//go:build e2e

package vlog_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/editorpost/donq/pkg/vlog"
	"log/slog"
	"testing"
	"time"
)

func TestNewElasticAPI(t *testing.T) {
	uri := "http://localhost:45191/insert/elasticsearch/"
	vlog.VictoriaLogger(uri, WindmillVars()...)
	slog.Info(
		"test message",
		slog.String("url", gofakeit.URL()),
	)
	time.Sleep(5000 * time.Millisecond)
}
