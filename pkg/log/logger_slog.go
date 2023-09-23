package log

import (
	"context"
	"log/slog"
	"os"
)

const (
	SlogLevelDebug   = slog.LevelDebug
	SlogLevelInfo    = slog.LevelInfo
	SlogLevelWarning = slog.LevelWarn
	SlogLevelError   = slog.LevelError
	SlogLevelFatal   = slog.Level(12)
	SlogLevelOff     = slog.Level(16)
)

var slogLevel = map[slog.Level]string{
	SlogLevelDebug:   "DEBUG",
	SlogLevelInfo:    "INFO",
	SlogLevelWarning: "WARN",
	SlogLevelError:   "ERROR",
	SlogLevelFatal:   "FATAL",
	SlogLevelOff:     "OFF",
}

var loggerLevel = map[LoggerLevel]slog.Level{
	UndefinedLoggerLevel: SlogLevelOff,
	DebugLoggerLevel:     SlogLevelDebug,
	InfoLoggerLevel:      SlogLevelInfo,
	WarnLoggerLevel:      SlogLevelWarning,
	ErrorLoggerLevel:     SlogLevelError,
	FatalLoggerLevel:     SlogLevelFatal,
	OffLoggerLevel:       SlogLevelOff,
}

type SlogLogger struct {
	internal *slog.Logger
}

func NewDefaultLogger(level LoggerLevel) *SlogLogger {
	opts := &slog.HandlerOptions{
		Level: loggerLevel[level],
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				a.Value = slog.StringValue(slogLevel[level])
			}
			return a
		},
	}
	internal := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(internal)
	slogLogger := &SlogLogger{internal: internal}
	singleton.Store(slogLogger)
	return slogLogger
}

func (logger *SlogLogger) Debug(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelDebug, msg, args...)
}

func (logger *SlogLogger) Info(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelInfo, msg, args...)
}

func (logger *SlogLogger) Warn(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelWarning, msg, args...)
}

func (logger *SlogLogger) Error(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelError, msg, args...)
}

func (logger *SlogLogger) Fatal(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelFatal, msg, args...)
}