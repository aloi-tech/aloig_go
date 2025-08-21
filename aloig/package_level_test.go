package aloig

import (
	"context"
	"errors"
	"os"
	"testing"
)

// TestPackageLevelContextFunctionsWork tests that context functions work without errors
func TestPackageLevelContextFunctionsWork(t *testing.T) {
	ctx := WithTraceID(context.Background(), "test-trace-123")
	ctx = WithRequestID(ctx, "test-request-456")

	// Test context functions - only verify they don't panic
	testCases := []struct {
		name    string
		logFunc func()
	}{
		{"DebugContext", func() { DebugContext(ctx, "test debug context") }},
		{"DebugfContext", func() { DebugfContext(ctx, "test debug context %s", "formatted") }},
		{"InfoContext", func() { InfoContext(ctx, "test info context") }},
		{"InfofContext", func() { InfofContext(ctx, "test info context %s", "formatted") }},
		{"WarnContext", func() { WarnContext(ctx, "test warn context") }},
		{"WarningContext", func() { WarningContext(ctx, "test warning context") }},
		{"WarnfContext", func() { WarnfContext(ctx, "test warn context %s", "formatted") }},
		{"WarningfContext", func() { WarningfContext(ctx, "test warning context %s", "formatted") }},
		{"ErrorContext", func() { ErrorContext(ctx, "test error context") }},
		{"ErrorfContext", func() { ErrorfContext(ctx, "test error context %s", "formatted") }},
		{"TraceContext", func() { TraceContext(ctx, "test trace context") }},
		{"TracefContext", func() { TracefContext(ctx, "test trace context %s", "formatted") }},
		{"PrintContext", func() { PrintContext(ctx, "test print context") }},
		{"PrintfContext", func() { PrintfContext(ctx, "test print context %s", "formatted") }},
		{"PrintlnContext", func() { PrintlnContext(ctx, "test println context") }},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Only verify that the function doesn't panic
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Function %s caused panic: %v", tc.name, r)
				}
			}()
			tc.logFunc()
		})
	}
}

// TestPackageLevelWithFieldsWork tests WithField, WithFields, WithError, WithContext functions
func TestPackageLevelWithFieldsWork(t *testing.T) {
	// Only verify that functions don't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("WithFields functions caused panic: %v", r)
		}
	}()

	// Test WithField
	WithField("key", "value").Info("test with field")

	// Test WithFields
	WithFields(map[string]interface{}{"key1": "value1", "key2": "value2"}).Info("test with fields")

	// Test WithError
	testError := errors.New("test error")
	WithError(testError).Info("test with error")

	// Test WithContext
	ctx := WithTraceID(context.Background(), "test-trace")
	WithContext(ctx).Info("test with context")
}

// TestPackageLevelWithNilContextWork tests behavior with nil context
func TestPackageLevelWithNilContextWork(t *testing.T) {
	// Only verify that functions don't panic with nil context
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Nil context functions caused panic: %v", r)
		}
	}()

	InfoContext(nil, "test with nil context")
}

// TestPackageLevelChainingWork tests chaining of package-level functions
func TestPackageLevelChainingWork(t *testing.T) {
	// Only verify that chaining doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Chaining caused panic: %v", r)
		}
	}()

	// Test chaining multiple methods
	WithField("field1", "value1").
		WithField("field2", "value2").
		WithError(errors.New("test error")).
		Info("test chained message")
}

// TestPackageLevelEnvironmentVariables tests that functions work with environment variables
func TestPackageLevelEnvironmentVariables(t *testing.T) {
	// Set environment variables for test
	os.Setenv("ENVIRONMENT", "test")
	os.Setenv("APP_NAME", "test-app")

	// Get configuration to verify environment variables are read
	config := DefaultConfig()

	// Check that environment variables are correctly read
	if config.Environment != "test" {
		t.Errorf("Expected Environment='test', got '%s'", config.Environment)
	}
	if config.AppName != "test-app" {
		t.Errorf("Expected AppName='test-app', got '%s'", config.AppName)
	}

	// Clean up
	os.Unsetenv("ENVIRONMENT")
	os.Unsetenv("APP_NAME")
}
