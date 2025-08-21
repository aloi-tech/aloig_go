package aloig

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

// TestCallerJSONFormatter tests that the custom JSON formatter includes caller information
func TestCallerJSONFormatter(t *testing.T) {
	formatter := &CallerJSONFormatter{JSONFormatter: &logrus.JSONFormatter{}}

	entry := &logrus.Entry{
		Message: "test message",
		Level:   logrus.InfoLevel,
		Caller: &runtime.Frame{
			File:     "/path/to/test.go",
			Line:     123,
			Function: "github.com/test.TestFunction",
		},
		Data: make(logrus.Fields),
	}

	output, err := formatter.Format(entry)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	outputStr := string(output)

	// Check that output contains caller information
	if !strings.Contains(outputStr, "test.go:123") {
		t.Error("Log output should contain 'test.go:123'")
	}
	if !strings.Contains(outputStr, "TestFunction") {
		t.Error("Log output should contain 'TestFunction'")
	}
}

// TestGetFunctionName tests function name extraction
func TestGetFunctionName(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"github.com/test.TestFunction", "TestFunction"},
		{"main.function", "function"},
		{"simple", "simple"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := getFunctionName(tc.input)
			if result != tc.expected {
				t.Errorf("Expected getFunctionName('%s') = '%s', got '%s'", tc.input, tc.expected, result)
			}
		})
	}
}

// TestGetLogLevelFromEnv tests getting log level from environment variables
func TestGetLogLevelFromEnv(t *testing.T) {
	testCases := []struct {
		envVar     string
		envValue   string
		defaultVal string
		expected   logrus.Level
	}{
		{"LOG_LEVEL", "debug", "info", logrus.DebugLevel},
		{"LOG_LEVEL", "info", "warn", logrus.InfoLevel},
		{"LOG_LEVEL", "warn", "error", logrus.WarnLevel},
		{"LOG_LEVEL", "error", "debug", logrus.ErrorLevel},
		{"LOG_LEVEL", "fatal", "info", logrus.FatalLevel},
		{"LOG_LEVEL", "panic", "info", logrus.PanicLevel},
		{"LOG_LEVEL", "trace", "info", logrus.TraceLevel},
		{"LOG_LEVEL", "invalid", "info", logrus.InfoLevel}, // Invalid should default
		{"LOG_LEVEL", "", "warn", logrus.WarnLevel},        // Empty should use default
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s=%s", tc.envVar, tc.envValue), func(t *testing.T) {
			// Set environment variable
			if tc.envValue != "" {
				os.Setenv(tc.envVar, tc.envValue)
			} else {
				os.Unsetenv(tc.envVar)
			}

			result := GetLogLevelFromEnv(tc.envVar, tc.defaultVal)

			if result != tc.expected {
				t.Errorf("Expected GetLogLevelFromEnv('%s', '%s') = %v, got %v", tc.envVar, tc.defaultVal, tc.expected, result)
			}

			// Clean up
			os.Unsetenv(tc.envVar)
		})
	}
}

// TestFlushSentry tests that FlushSentry works correctly
func TestFlushSentry(t *testing.T) {
	// Test in non-Sentry environment
	os.Setenv("ENVIRONMENT", "dev")
	FlushSentry() // Should not panic

	// Test in Sentry environment without DSN
	os.Setenv("ENVIRONMENT", "prod")
	os.Unsetenv("SENTRY_DSN")
	FlushSentry() // Should not panic

	// Test in Sentry environment with DSN
	os.Setenv("ENVIRONMENT", "prod")
	os.Setenv("SENTRY_DSN", "https://test@test.ingest.sentry.io/test")
	FlushSentry() // Should not panic

	// Clean up
	os.Unsetenv("ENVIRONMENT")
	os.Unsetenv("SENTRY_DSN")
}

// TestAloigContextFunctionsWork tests that context functions work without errors
func TestAloigContextFunctionsWork(t *testing.T) {
	ctx := context.Background()
	ctx = WithTraceID(ctx, "test-trace-123")
	ctx = WithRequestID(ctx, "test-request-456")

	// Only verify that functions don't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Context functions caused panic: %v", r)
		}
	}()

	DebugContext(ctx, "test debug context")
	InfoContext(ctx, "test info context")
	ErrorContext(ctx, "test error context")
}

// TestAloigWithNilContextWork tests behavior with nil context
func TestAloigWithNilContextWork(t *testing.T) {
	// Only verify that functions don't panic with nil context
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Nil context functions caused panic: %v", r)
		}
	}()

	InfoContext(nil, "test with nil context")
}

// TestAloigWithEmptyContextWork tests behavior with empty context
func TestAloigWithEmptyContextWork(t *testing.T) {
	// Only verify that functions don't panic with empty context
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Empty context functions caused panic: %v", r)
		}
	}()

	ctx := context.Background()
	InfoContext(ctx, "test with empty context")
}
