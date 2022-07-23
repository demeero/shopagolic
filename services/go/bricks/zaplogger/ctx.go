package zaplogger

import (
	"context"

	"go.uber.org/zap"
)

type logCtxType struct{}

var logCtxKey = logCtxType{}

// FromCtx returns a logger set on the context, or the global zap.L() if none is found.
func FromCtx(ctx context.Context) *zap.Logger {
	return fromCtxOrDefault(ctx, zap.L())
}

// fromCtxOrDefault returns a logger set on the context, or the caller specified default logger.
func fromCtxOrDefault(ctx context.Context, def *zap.Logger) *zap.Logger {
	logger, ok := ctx.Value(logCtxKey).(*zap.Logger)
	if !ok || logger == nil {
		return def
	}
	return logger
}

// ToCtx add logger into context.
func ToCtx(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, logCtxKey, logger)
}
