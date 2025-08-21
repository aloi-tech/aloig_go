package aloig

import (
	"context"
	"strings"
	"testing"
)

// TestWithTraceID tests that WithTraceID correctly adds trace ID to context
func TestWithTraceID(t *testing.T) {
	ctx := context.Background()
	traceID := "test-trace-123"

	ctxWithTrace := WithTraceID(ctx, traceID)

	result := GetTraceID(ctxWithTrace)
	if result != traceID {
		t.Errorf("Expected trace ID '%s', got '%s'", traceID, result)
	}
}

// TestGetTraceID tests that GetTraceID correctly retrieves trace ID from context
func TestGetTraceID(t *testing.T) {
	testCases := []struct {
		name     string
		ctx      context.Context
		expected string
	}{
		{
			"Context with trace ID",
			WithTraceID(context.Background(), "test-trace-123"),
			"test-trace-123",
		},
		{
			"Context without trace ID",
			context.Background(),
			"",
		},
		{
			"Nil context",
			nil,
			"",
		},
		{
			"Context with empty trace ID",
			WithTraceID(context.Background(), ""),
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetTraceID(tc.ctx)
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}

// TestEnsureTraceID tests that EnsureTraceID works correctly
func TestEnsureTraceID(t *testing.T) {
	testCases := []struct {
		name           string
		ctx            context.Context
		shouldGenerate bool
	}{
		{
			"Context without trace ID - should generate new one",
			context.Background(),
			true,
		},
		{
			"Context with existing trace ID - should keep existing one",
			WithTraceID(context.Background(), "existing-trace"),
			false,
		},
		{
			"Nil context - should generate new one",
			nil,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultCtx, resultTraceID := EnsureTraceID(tc.ctx)

			if resultCtx == nil {
				t.Error("Resulting context should not be nil")
			}

			if resultTraceID == "" {
				t.Error("Resulting trace ID should not be empty")
			}

			// Verify that trace ID is in context
			ctxTraceID := GetTraceID(resultCtx)
			if ctxTraceID != resultTraceID {
				t.Errorf("Trace ID in context '%s' does not match returned '%s'", ctxTraceID, resultTraceID)
			}

			// If should not generate, verify it's the same as original
			if !tc.shouldGenerate {
				originalTraceID := GetTraceID(tc.ctx)
				if resultTraceID != originalTraceID {
					t.Errorf("Should keep original trace ID '%s', but got '%s'", originalTraceID, resultTraceID)
				}
			}
		})
	}
}

// TestGenerateTraceID tests that GenerateTraceID generates valid trace IDs
func TestGenerateTraceID(t *testing.T) {
	// Generate multiple trace IDs to ensure they're different
	traceID1 := GenerateTraceID()
	traceID2 := GenerateTraceID()

	if traceID1 == traceID2 {
		t.Error("Generated trace IDs should be different")
	}

	if len(traceID1) == 0 {
		t.Error("Generated trace ID should not be empty")
	}

	if len(traceID2) == 0 {
		t.Error("Generated trace ID should not be empty")
	}

	// Verify format (should be UUID without dashes)
	if strings.Contains(traceID1, "-") {
		t.Error("Generated trace ID should not contain dashes")
	}

	if strings.Contains(traceID2, "-") {
		t.Error("Generated trace ID should not contain dashes")
	}
}

// TestWithRequestID tests that WithRequestID correctly adds request ID to context
func TestWithRequestID(t *testing.T) {
	ctx := context.Background()
	requestID := "test-request-456"

	ctxWithRequest := WithRequestID(ctx, requestID)

	result := GetRequestID(ctxWithRequest)
	if result != requestID {
		t.Errorf("Expected request ID '%s', got '%s'", requestID, result)
	}
}

// TestGetRequestID tests that GetRequestID correctly retrieves request ID from context
func TestGetRequestID(t *testing.T) {
	testCases := []struct {
		name     string
		ctx      context.Context
		expected string
	}{
		{
			"Context with request ID",
			WithRequestID(context.Background(), "test-request-456"),
			"test-request-456",
		},
		{
			"Context without request ID",
			context.Background(),
			"",
		},
		{
			"Nil context",
			nil,
			"",
		},
		{
			"Context with empty request ID",
			WithRequestID(context.Background(), ""),
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetRequestID(tc.ctx)
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}

// TestWithUserID tests that WithUserID correctly adds user ID to context
func TestWithUserID(t *testing.T) {
	ctx := context.Background()
	userID := "test-user-789"

	ctxWithUser := WithUserID(ctx, userID)

	result := GetUserID(ctxWithUser)
	if result != userID {
		t.Errorf("Expected user ID '%s', got '%s'", userID, result)
	}
}

// TestGetUserID tests that GetUserID correctly retrieves user ID from context
func TestGetUserID(t *testing.T) {
	testCases := []struct {
		name     string
		ctx      context.Context
		expected string
	}{
		{
			"Context with user ID",
			WithUserID(context.Background(), "test-user-789"),
			"test-user-789",
		},
		{
			"Context without user ID",
			context.Background(),
			"",
		},
		{
			"Nil context",
			nil,
			"",
		},
		{
			"Context with empty user ID",
			WithUserID(context.Background(), ""),
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetUserID(tc.ctx)
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}

// TestWithSessionID tests that WithSessionID correctly adds session ID to context
func TestWithSessionID(t *testing.T) {
	ctx := context.Background()
	sessionID := "test-session-abc"

	ctxWithSession := WithSessionID(ctx, sessionID)

	result := GetSessionID(ctxWithSession)
	if result != sessionID {
		t.Errorf("Expected session ID '%s', got '%s'", sessionID, result)
	}
}

// TestGetSessionID tests that GetSessionID correctly retrieves session ID from context
func TestGetSessionID(t *testing.T) {
	testCases := []struct {
		name     string
		ctx      context.Context
		expected string
	}{
		{
			"Context with session ID",
			WithSessionID(context.Background(), "test-session-abc"),
			"test-session-abc",
		},
		{
			"Context without session ID",
			context.Background(),
			"",
		},
		{
			"Nil context",
			nil,
			"",
		},
		{
			"Context with empty session ID",
			WithSessionID(context.Background(), ""),
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetSessionID(tc.ctx)
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}

// TestExtractContextFields tests that ExtractContextFields correctly extracts all context fields
func TestExtractContextFields(t *testing.T) {
	// Create context with all fields
	ctx := context.Background()
	ctx = WithTraceID(ctx, "test-trace-123")
	ctx = WithRequestID(ctx, "test-request-456")
	ctx = WithUserID(ctx, "test-user-789")
	ctx = WithSessionID(ctx, "test-session-abc")

	fields := ExtractContextFields(ctx)

	expectedFields := map[string]interface{}{
		"trace_id":   "test-trace-123",
		"request_id": "test-request-456",
		"user_id":    "test-user-789",
		"session_id": "test-session-abc",
	}

	if len(fields) != len(expectedFields) {
		t.Errorf("Expected %d fields, got %d", len(expectedFields), len(fields))
	}

	for key, expectedValue := range expectedFields {
		if value, exists := fields[key]; !exists {
			t.Errorf("Expected field '%s' to exist", key)
		} else if value != expectedValue {
			t.Errorf("Expected field '%s' to be '%v', got '%v'", key, expectedValue, value)
		}
	}
}

// TestExtractContextFieldsEmpty tests that ExtractContextFields handles empty context correctly
func TestExtractContextFieldsEmpty(t *testing.T) {
	// Test with empty context
	ctx := context.Background()
	fields := ExtractContextFields(ctx)

	if len(fields) != 0 {
		t.Errorf("Expected 0 fields for empty context, got %d", len(fields))
	}

	// Test with nil context
	fields = ExtractContextFields(nil)

	if len(fields) != 0 {
		t.Errorf("Expected 0 fields for nil context, got %d", len(fields))
	}
}

// TestExtractContextFieldsPartial tests that ExtractContextFields handles partial context correctly
func TestExtractContextFieldsPartial(t *testing.T) {
	// Test with only trace ID
	ctx := WithTraceID(context.Background(), "test-trace-123")
	fields := ExtractContextFields(ctx)

	if len(fields) != 1 {
		t.Errorf("Expected 1 field, got %d", len(fields))
	}

	if fields["trace_id"] != "test-trace-123" {
		t.Errorf("Expected trace_id to be 'test-trace-123', got '%v'", fields["trace_id"])
	}

	// Test with only request ID
	ctx = WithRequestID(context.Background(), "test-request-456")
	fields = ExtractContextFields(ctx)

	if len(fields) != 1 {
		t.Errorf("Expected 1 field, got %d", len(fields))
	}

	if fields["request_id"] != "test-request-456" {
		t.Errorf("Expected request_id to be 'test-request-456', got '%v'", fields["request_id"])
	}
}

// TestContextChaining tests that context functions can be chained correctly
func TestContextChaining(t *testing.T) {
	ctx := context.Background()
	ctx = WithTraceID(ctx, "trace-123")
	ctx = WithRequestID(ctx, "request-456")
	ctx = WithUserID(ctx, "user-789")
	ctx = WithSessionID(ctx, "session-abc")

	// Verify all fields are present
	if GetTraceID(ctx) != "trace-123" {
		t.Error("Trace ID not preserved in chaining")
	}

	if GetRequestID(ctx) != "request-456" {
		t.Error("Request ID not preserved in chaining")
	}

	if GetUserID(ctx) != "user-789" {
		t.Error("User ID not preserved in chaining")
	}

	if GetSessionID(ctx) != "session-abc" {
		t.Error("Session ID not preserved in chaining")
	}

	// Verify extraction works correctly
	fields := ExtractContextFields(ctx)
	expectedCount := 4
	if len(fields) != expectedCount {
		t.Errorf("Expected %d fields after chaining, got %d", expectedCount, len(fields))
	}
}
