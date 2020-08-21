package testutil

import "testing"

func NewLoggerAdapter(tb testing.TB) LoggerAdapter {
	return LoggerAdapter{tb}
}

type LoggerAdapter struct {
	tb testing.TB
}

func (l LoggerAdapter) Printf(format string, args ...interface{}) {
	l.tb.Logf(format, args...)
}
