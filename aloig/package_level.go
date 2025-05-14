package aloig

import (
	"context"
)

// Este archivo contiene funciones de conveniencia a nivel de paquete
// para permitir el uso de aloig.Info(), aloig.Error(), etc. directamente

// Debug registra un mensaje de nivel debug usando el logger singleton
func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

// Debugf registra un mensaje de nivel debug con formato usando el logger singleton
func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

// Info registra un mensaje de nivel info usando el logger singleton
func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

// Infof registra un mensaje de nivel info con formato usando el logger singleton
func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

// Warn registra un mensaje de nivel warning usando el logger singleton
func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

// Warning es un alias de Warn que registra un mensaje de nivel warning
func Warning(args ...interface{}) {
	GetLogger().Warning(args...)
}

// Warnf registra un mensaje de nivel warning con formato usando el logger singleton
func Warnf(format string, args ...interface{}) {
	GetLogger().Warnf(format, args...)
}

// Warningf es un alias de Warnf que registra un mensaje de nivel warning con formato
func Warningf(format string, args ...interface{}) {
	GetLogger().Warningf(format, args...)
}

// Error registra un mensaje de nivel error usando el logger singleton
func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

// Errorf registra un mensaje de nivel error con formato usando el logger singleton
func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

// Fatal registra un mensaje de nivel fatal usando el logger singleton
// y luego hace que la aplicación termine con un status code no-cero
func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

// Fatalf registra un mensaje de nivel fatal con formato usando el logger singleton
// y luego hace que la aplicación termine con un status code no-cero
func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

// Panic registra un mensaje de nivel panic usando el logger singleton
// y luego lanza un panic con el mensaje formateado
func Panic(args ...interface{}) {
	GetLogger().Panic(args...)
}

// Panicf registra un mensaje de nivel panic con formato usando el logger singleton
// y luego lanza un panic con el mensaje formateado
func Panicf(format string, args ...interface{}) {
	GetLogger().Panicf(format, args...)
}

// Printf imprime un mensaje formateado usando el logger singleton
func Printf(format string, args ...interface{}) {
	GetLogger().Printf(format, args...)
}

// Print imprime un mensaje usando el logger singleton
func Print(args ...interface{}) {
	GetLogger().Print(args...)
}

// Println imprime un mensaje con nueva línea usando el logger singleton
func Println(args ...interface{}) {
	GetLogger().Println(args...)
}

// Trace registra un mensaje de nivel trace usando el logger singleton
func Trace(args ...interface{}) {
	GetLogger().Trace(args...)
}

// Tracef registra un mensaje de nivel trace con formato usando el logger singleton
func Tracef(format string, args ...interface{}) {
	GetLogger().Tracef(format, args...)
}

// WithField retorna una nueva entrada de log con el campo key=value añadido
func WithField(key string, value interface{}) Logger {
	return GetLogger().WithField(key, value)
}

// WithFields retorna una nueva entrada de log con los campos añadidos
func WithFields(fields map[string]interface{}) Logger {
	return GetLogger().WithFields(fields)
}

// WithError retorna una nueva entrada de log con un error añadido
func WithError(err error) Logger {
	return GetLogger().WithError(err)
}

// WithContext retorna una nueva entrada de log con el contexto añadido
func WithContext(ctx context.Context) Logger {
	return GetLogger().WithContext(ctx)
}

// DebugContext registra un mensaje de debug usando el contexto dado
func DebugContext(ctx context.Context, args ...interface{}) {
	GetLogger().DebugContext(ctx, args...)
}

// DebugfContext registra un mensaje de debug formateado usando el contexto dado
func DebugfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().DebugfContext(ctx, format, args...)
}

// InfoContext registra un mensaje de info usando el contexto dado
func InfoContext(ctx context.Context, args ...interface{}) {
	GetLogger().InfoContext(ctx, args...)
}

// InfofContext registra un mensaje de info formateado usando el contexto dado
func InfofContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().InfofContext(ctx, format, args...)
}

// WarnContext registra un mensaje de advertencia usando el contexto dado
func WarnContext(ctx context.Context, args ...interface{}) {
	GetLogger().WarnContext(ctx, args...)
}

// WarnfContext registra un mensaje de advertencia formateado usando el contexto dado
func WarnfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().WarnfContext(ctx, format, args...)
}

// WarningContext registra un mensaje de advertencia usando el contexto dado
func WarningContext(ctx context.Context, args ...interface{}) {
	GetLogger().WarningContext(ctx, args...)
}

// WarningfContext registra un mensaje de advertencia formateado usando el contexto dado
func WarningfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().WarningfContext(ctx, format, args...)
}

// ErrorContext registra un mensaje de error usando el contexto dado
func ErrorContext(ctx context.Context, args ...interface{}) {
	GetLogger().ErrorContext(ctx, args...)
}

// ErrorfContext registra un mensaje de error formateado usando el contexto dado
func ErrorfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().ErrorfContext(ctx, format, args...)
}

// FatalContext registra un mensaje fatal usando el contexto dado
func FatalContext(ctx context.Context, args ...interface{}) {
	GetLogger().FatalContext(ctx, args...)
}

// FatalfContext registra un mensaje fatal formateado usando el contexto dado
func FatalfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().FatalfContext(ctx, format, args...)
}

// PanicContext registra un mensaje de pánico usando el contexto dado
func PanicContext(ctx context.Context, args ...interface{}) {
	GetLogger().PanicContext(ctx, args...)
}

// PanicfContext registra un mensaje de pánico formateado usando el contexto dado
func PanicfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().PanicfContext(ctx, format, args...)
}

// PrintContext imprime un mensaje usando el contexto dado
func PrintContext(ctx context.Context, args ...interface{}) {
	GetLogger().PrintContext(ctx, args...)
}

// PrintfContext imprime un mensaje formateado usando el contexto dado
func PrintfContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().PrintfContext(ctx, format, args...)
}

// PrintlnContext imprime un mensaje con salto de línea usando el contexto dado
func PrintlnContext(ctx context.Context, args ...interface{}) {
	GetLogger().PrintlnContext(ctx, args...)
}

// TraceContext registra un mensaje de trace usando el contexto dado
func TraceContext(ctx context.Context, args ...interface{}) {
	GetLogger().TraceContext(ctx, args...)
}

// TracefContext registra un mensaje de trace formateado usando el contexto dado
func TracefContext(ctx context.Context, format string, args ...interface{}) {
	GetLogger().TracefContext(ctx, format, args...)
}
