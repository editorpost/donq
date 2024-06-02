package vlog

import (
	"log/slog"
	"strings"
)

// StdoutLogger sets the vlog as default slog logger
func StdoutLogger(lvl slog.Level, attrs ...slog.Attr) {
	slog.SetDefault(New(lvl, attrs...))
}

// VictoriaLogger sets the vlog as default slog logger with VictoriaMetrics sender
func VictoriaLogger(uri string, lvl slog.Level, attrs ...slog.Attr) {
	slog.SetDefault(NewVictoriaLogs(uri, lvl, attrs...))
}

// New creates a new vlog with Stdout sender.
func New(lvl slog.Level, attrs ...slog.Attr) *slog.Logger {

	// buffering/sending logs
	pool := NewPool(StdoutSender(Mapper))
	go pool.Ticker(1)

	// catching logs
	handler := NewBaseHandler(slog.LevelInfo, pool, attrs)
	return slog.New(handler)
}

// NewVictoriaLogs creates a new vlog with VictoriaMetrics sender.
func NewVictoriaLogs(uri string, lvl slog.Level, attrs ...slog.Attr) *slog.Logger {

	uri, _ = strings.CutPrefix(uri, "/")
	uri += "/insert/elasticsearch/"

	return NewElastic(uri, lvl, attrs...)
}

// NewElastic creates a new vlog with ElasticSearch sender.
func NewElastic(uri string, lvl slog.Level, attrs ...slog.Attr) *slog.Logger {

	ingester, err := NewElasticIngest(uri, Mapper)
	if err != nil {
		// fallback to stdout
		return New(lvl, attrs...)
	}
	// buffering/sending logs
	pool := NewPool(ingester.Sender())
	go pool.Ticker(5)

	// catching logs
	handler := NewBaseHandler(slog.LevelInfo, pool, attrs)
	return slog.New(handler)
}
