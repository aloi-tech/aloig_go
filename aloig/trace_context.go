package aloig

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

type contextKey string

const (
	// TraceIDKey is the key used for trace ID in context
	TraceIDKey contextKey = "trace_id"

	// RequestIDKey is the key used for request ID in context
	RequestIDKey contextKey = "request_id"

	// UserIDKey is the key used for user ID in context
	UserIDKey contextKey = "user_id"

	// SessionIDKey is the key used for session ID in context
	SessionIDKey contextKey = "session_id"
)

// WithTraceID returns a new context with the specified trace ID
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

// GetTraceID gets the trace ID from context
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

// EnsureTraceID ensures there's a trace ID in the context
// If it doesn't exist, creates a new one
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

// GenerateTraceID generates a new random trace ID
func GenerateTraceID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// WithRequestID returns a new context with the specified request ID
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

// GetRequestID gets the request ID from context
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

// WithUserID returns a new context with the specified user ID
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// GetUserID gets the user ID from context
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

// WithSessionID returns a new context with the specified session ID
func WithSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, SessionIDKey, sessionID)
}

// GetSessionID gets the session ID from context
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

// ExtractContextFields extracts all context fields into a map
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
