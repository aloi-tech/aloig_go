package aloig

import (
	"context"
)

// This file contains package-level convenience functions
// to allow using aloig.Info(), aloig.Error(), etc. directly

// Debug logs a debug level message using the singleton logger
func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

// Debugf logs a formatted debug level message using the singleton logger
func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

// Info logs an info level message using the singleton logger
func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

// Infof logs a formatted info level message using the singleton logger
func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

// Warn logs a warning level message using the singleton logger
func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

// Warning is an alias for Warn that logs a warning level message
func Warning(args ...interface{}) {
	GetLogger().Warning(args...)
}

// Warnf logs a formatted warning level message using the singleton logger
func Warnf(format string, args ...interface{}) {
	GetLogger().Warnf(format, args...)
}

// Warningf is an alias for Warnf that logs a formatted warning level message
func Warningf(format string, args ...interface{}) {
	GetLogger().Warningf(format, args...)
}

// Error logs an error level message using the singleton logger
func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

// Errorf logs a formatted error level message using the singleton logger
func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

// Fatal logs a fatal level message using the singleton logger
// and then makes the application exit with a non-zero status code
func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

// Fatalf logs a formatted fatal level message using the singleton logger
// and then makes the application exit with a non-zero status code
func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

// Panic logs a panic level message using the singleton logger
// and then throws a panic with the formatted message
func Panic(args ...interface{}) {
	GetLogger().Panic(args...)
}

// Panicf logs a formatted panic level message using the singleton logger
// and then throws a panic with the formatted message
func Panicf(format string, args ...interface{}) {
	GetLogger().Panicf(format, args...)
}

// Printf prints a formatted message using the singleton logger
func Printf(format string, args ...interface{}) {
	GetLogger().Printf(format, args...)
}

// Print prints a message using the singleton logger
func Print(args ...interface{}) {
	GetLogger().Print(args...)
}

// Println prints a message with newline using the singleton logger
func Println(args ...interface{}) {
	GetLogger().Println(args...)
}

// Trace logs a trace level message using the singleton logger
func Trace(args ...interface{}) {
	GetLogger().Trace(args...)
}

// Tracef logs a formatted trace level message using the singleton logger
func Tracef(format string, args ...interface{}) {
	GetLogger().Tracef(format, args...)
}

// WithField returns a new log entry with the key=value field added
func WithField(key string, value interface{}) Logger {
	return GetLogger().WithField(key, value)
}

// WithFields returns a new log entry with the fields added
func WithFields(fields map[string]interface{}) Logger {
	return GetLogger().WithFields(fields)
}

// WithError returns a new log entry with an error added
func WithError(err error) Logger {
	return GetLogger().WithError(err)
}

// WithContext returns a new log entry with the context added
func WithContext(ctx context.Context) Logger {
	return GetLogger().WithContext(ctx)
}

// DebugContext logs a debug message using the given context
func DebugContext(ctx context.Context, args ...interface{}) {
	GetLogger().DebugContext(ctx, args...)
}

// DebugfContext logs a formatted debug message using the given context
func DebugfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().DebugfContext(ctx, format, args...)
}

// InfoContext logs an info message using the given context
func InfoContext(ctx context.Context, args ...interface{}) {
	GetLogger().InfoContext(ctx, args...)
}

// InfofContext logs a formatted info message using the given context
func InfofContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().InfofContext(ctx, format, args...)
}

// WarnContext logs a warning message using the given context
func WarnContext(ctx context.Context, args ...interface{}) {
	GetLogger().WarnContext(ctx, args...)
}

// WarnfContext logs a formatted warning message using the given context
func WarnfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().WarnfContext(ctx, format, args...)
}

// WarningContext logs a warning message using the given context
func WarningContext(ctx context.Context, args ...interface{}) {
	GetLogger().WarningContext(ctx, args...)
}

// WarningfContext logs a formatted warning message using the given context
func WarningfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().WarningfContext(ctx, format, args...)
}

// ErrorContext logs an error message using the given context
func ErrorContext(ctx context.Context, args ...interface{}) {
	GetLogger().ErrorContext(ctx, args...)
}

// ErrorfContext logs a formatted error message using the given context
func ErrorfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().ErrorfContext(ctx, format, args...)
}

// FatalContext logs a fatal message using the given context
func FatalContext(ctx context.Context, args ...interface{}) {
	GetLogger().FatalContext(ctx, args...)
}

// FatalfContext logs a formatted fatal message using the given context
func FatalfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().FatalfContext(ctx, format, args...)
}

// PanicContext logs a panic message using the given context
func PanicContext(ctx context.Context, args ...interface{}) {
	GetLogger().PanicContext(ctx, args...)
}

// PanicfContext logs a formatted panic message using the given context
func PanicfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().PanicfContext(ctx, format, args...)
}

// PrintContext prints a message using the given context
func PrintContext(ctx context.Context, args ...interface{}) {
	GetLogger().PrintContext(ctx, args...)
}

// PrintfContext prints a formatted message using the given context
func PrintfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().PrintfContext(ctx, format, args...)
}

// PrintlnContext prints a message with newline using the given context
func PrintlnContext(ctx context.Context, args ...interface{}) {
	GetLogger().PrintlnContext(ctx, args...)
}

// TraceContext logs a trace message using the given context
func TraceContext(ctx context.Context, args ...interface{}) {
	GetLogger().TraceContext(ctx, args...)
}

// TracefContext logs a formatted trace message using the given context
func TracefContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().TracefContext(ctx, format, args...)
}
