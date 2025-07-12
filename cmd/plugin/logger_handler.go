package plugin

import (
	"context"
	"log/slog"
)

// NOTE: This file is just skeleton. Read ref below
// ref. https://github.com/golang/example/blob/master/slog-handler-guide/README.md#the-withattrs-method
type LoggerHandler interface {
	Enabled(context.Context, slog.Leveler) bool
	Handle(context.Context, slog.Record) error
	WithAttrs(attrs []slog.Attr) slog.Handler
	WithGroup(name string) slog.Handler
}
