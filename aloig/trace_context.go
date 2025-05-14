package aloig

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

type contextKey string

const (
	// TraceIDKey es la clave usada para el trace ID en el contexto
	TraceIDKey contextKey = "trace_id"

	// RequestIDKey es la clave usada para el ID de la petición en el contexto
	RequestIDKey contextKey = "request_id"

	// UserIDKey es la clave usada para el ID del usuario en el contexto
	UserIDKey contextKey = "user_id"

	// SessionIDKey es la clave usada para el ID de sesión en el contexto
	SessionIDKey contextKey = "session_id"
)

// WithTraceID retorna un nuevo contexto con el trace ID especificado
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

// GetTraceID obtiene el trace ID del contexto
func GetTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	traceID, ok := ctx.Value(TraceIDKey).(string)
	if !ok || traceID == "" {
		return ""
	}
	return traceID
}

// EnsureTraceID asegura que haya un trace ID en el contexto
// Si no existe, crea uno nuevo
func EnsureTraceID(ctx context.Context) (context.Context, string) {
	if ctx == nil {
		ctx = context.Background()
	}

	traceID := GetTraceID(ctx)
	if traceID == "" {
		traceID = GenerateTraceID()
		ctx = WithTraceID(ctx, traceID)
	}

	return ctx, traceID
}

// GenerateTraceID genera un nuevo trace ID aleatorio
func GenerateTraceID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// WithRequestID retorna un nuevo contexto con el ID de petición especificado
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

// GetRequestID obtiene el ID de petición del contexto
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	requestID, ok := ctx.Value(RequestIDKey).(string)
	if !ok || requestID == "" {
		return ""
	}
	return requestID
}

// WithUserID retorna un nuevo contexto con el ID de usuario especificado
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// GetUserID obtiene el ID de usuario del contexto
func GetUserID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok || userID == "" {
		return ""
	}
	return userID
}

// WithSessionID retorna un nuevo contexto con el ID de sesión especificado
func WithSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, SessionIDKey, sessionID)
}

// GetSessionID obtiene el ID de sesión del contexto
func GetSessionID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	sessionID, ok := ctx.Value(SessionIDKey).(string)
	if !ok || sessionID == "" {
		return ""
	}
	return sessionID
}

// ExtractContextFields extrae todos los campos de contexto en un mapa
func ExtractContextFields(ctx context.Context) map[string]interface{} {
	fields := make(map[string]interface{})

	if traceID := GetTraceID(ctx); traceID != "" {
		fields["trace_id"] = traceID
	}

	if requestID := GetRequestID(ctx); requestID != "" {
		fields["request_id"] = requestID
	}

	if userID := GetUserID(ctx); userID != "" {
		fields["user_id"] = userID
	}

	if sessionID := GetSessionID(ctx); sessionID != "" {
		fields["session_id"] = sessionID
	}

	return fields
}
