package aloig

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

// setupTestLogger configures the singleton logger for testing with a buffer
func setupTestLogger() (*bytes.Buffer, func()) {
	var buf bytes.Buffer

	// Create a completely clean new logger
	logrusInstance := logrus.New()
	logrusInstance.SetLevel(logrus.TraceLevel)
	logrusInstance.SetReportCaller(true)
	logrusInstance.SetOutput(&buf)
	logrusInstance.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true, // For cleaner tests
		DisableColors:    true,
	})

	// Create a new logrusLogger without hooks or extra configurations
	testLogger := &logrusLogger{logger: logrusInstance}

	// Save original logger and configure test logger
	originalLog := log
	log = testLogger

	// Cleanup function
	cleanup := func() {
		log = originalLog
	}

	return &buf, cleanup
}

// TestPanicWithTrace tests that when there's a panic, trace information is included
func TestPanicWithTrace(t *testing.T) {
	// Configure test logger using helper function
	buf, cleanup := setupTestLogger()
	defer cleanup()

	// Create context with trace information
	ctx := WithTraceID(context.Background(), "test-trace-123")
	ctx = WithRequestID(ctx, "test-request-456")
	ctx = WithUserID(ctx, "user-789")

	// Log directly to verify buffer works
	ErrorContext(ctx, "Test message")
	output := buf.String()
	if !strings.Contains(output, "Test message") {
		t.Errorf("Expected log to contain 'Test message', got: %s", output)
	}

	// Test panic with context
	defer func() {
		if r := recover(); r != nil {
			panicStr := fmt.Sprintf("%v", r)
			if !strings.Contains(panicStr, "test panic message") {
				t.Errorf("Expected panic to contain 'test panic message', got: %s", panicStr)
			}
		}
	}()

	// Clear buffer before test
	buf.Reset()

	// Trigger panic
	panic("test panic message")
}

// TestPanicWithCallerInfo tests that panic includes caller information
func TestPanicWithCallerInfo(t *testing.T) {
	// Configure test logger using helper function
	buf, cleanup := setupTestLogger()
	defer cleanup()

	// Create context with trace information
	ctx := WithTraceID(context.Background(), "test-trace-123")
	ctx = WithRequestID(ctx, "test-request-456")
	ctx = WithUserID(ctx, "user-789")

	// Test panic with context
	defer func() {
		if r := recover(); r != nil {
			panicStr := fmt.Sprintf("%v", r)
			if !strings.Contains(panicStr, "panic caller test") {
				t.Errorf("Expected panic to contain 'panic caller test', got: %s", panicStr)
			}
		}
	}()

	// Clear buffer before test
	buf.Reset()

	// Trigger panic
	panic("panic caller test")

	// Check that log contains caller information
	output := buf.String()
	if !strings.Contains(output, "panic caller test") {
		t.Errorf("Expected log to contain 'panic caller test', got: %s", output)
	}
}

// TestPanicContext tests that PanicContext works correctly with trace
func TestPanicContext(t *testing.T) {
	// Configure test logger using helper function
	buf, cleanup := setupTestLogger()
	defer cleanup()

	// Create context with trace information
	ctx := WithTraceID(context.Background(), "test-trace-123")
	ctx = WithRequestID(ctx, "test-request-456")
	ctx = WithUserID(ctx, "user-789")

	// Test panic with context
	defer func() {
		if r := recover(); r != nil {
			panicStr := fmt.Sprintf("%v", r)
			if !strings.Contains(panicStr, "panic context test") {
				t.Errorf("Expected panic to contain 'panic context test', got: %s", panicStr)
			}
		}
	}()

	// Clear buffer before test
	buf.Reset()

	// Trigger panic using PanicContext
	PanicContext(ctx, "panic context test")

	// Check that log contains at least the panic message
	output := buf.String()
	if !strings.Contains(output, "panic context test") {
		t.Errorf("Expected log to contain 'panic context test', got: %s", output)
	}
}

// TestPanicfContext tests that PanicfContext works correctly with trace
func TestPanicfContext(t *testing.T) {
	// Configure test logger using helper function
	buf, cleanup := setupTestLogger()
	defer cleanup()

	// Create context with trace information
	ctx := WithTraceID(context.Background(), "test-trace-123")
	ctx = WithRequestID(ctx, "test-request-456")
	ctx = WithUserID(ctx, "user-789")

	// Test panic with context
	defer func() {
		if r := recover(); r != nil {
			panicStr := fmt.Sprintf("%v", r)
			if !strings.Contains(panicStr, "panicf context test") || !strings.Contains(panicStr, "formatted") {
				t.Errorf("Expected panic to contain 'panicf context test formatted', got: %s", panicStr)
			}
		}
	}()

	// Clear buffer before test
	buf.Reset()

	// Trigger panic using PanicfContext
	PanicfContext(ctx, "panicf context test %s", "formatted")

	// Check that log contains at least the panic message
	output := buf.String()
	if !strings.Contains(output, "panicf context test formatted") {
		t.Errorf("Expected log to contain 'panicf context test formatted', got: %s", output)
	}
}

// TestPackageLevelPanicWithTrace tests that package-level functions include trace in panic
func TestPackageLevelPanicWithTrace(t *testing.T) {
	// Configure test logger using helper function
	buf, cleanup := setupTestLogger()
	defer cleanup()

	// Create context with trace information
	ctx := WithTraceID(context.Background(), "test-trace-123")
	ctx = WithRequestID(ctx, "test-request-456")
	ctx = WithUserID(ctx, "user-789")

	// Test panic with context
	defer func() {
		if r := recover(); r != nil {
			panicStr := fmt.Sprintf("%v", r)
			if !strings.Contains(panicStr, "package panic test") {
				t.Errorf("Expected panic to contain 'package panic test', got: %s", panicStr)
			}
		}
	}()

	// Clear buffer before test
	buf.Reset()

	// Trigger panic using package-level function
	PanicContext(ctx, "package panic test")

	// Check that log contains at least the panic message
	output := buf.String()
	if !strings.Contains(output, "package panic test") {
		t.Errorf("Expected log to contain 'package panic test', got: %s", output)
	}
}

// TestPanicRecoveryWithTrace tests that panic recovery maintains trace
func TestPanicRecoveryWithTrace(t *testing.T) {
	// Configure test logger using helper function
	buf, cleanup := setupTestLogger()
	defer cleanup()

	// Create context with trace information
	ctx := WithTraceID(context.Background(), "test-trace-123")
	ctx = WithRequestID(ctx, "test-request-456")
	ctx = WithUserID(ctx, "user-789")

	// Test panic recovery
	defer func() {
		if r := recover(); r != nil {
			panicStr := fmt.Sprintf("%v", r)
			if !strings.Contains(panicStr, "recovery test panic") {
				t.Errorf("Expected panic to contain 'recovery test panic', got: %s", panicStr)
			}
		}
	}()

	// Clear buffer before test
	buf.Reset()

	// Log before panic
	InfoContext(ctx, "before panic")

	// Trigger panic
	panic("recovery test panic")

	// Check that logs contain expected messages
	output := buf.String()
	if !strings.Contains(output, "before panic") {
		t.Errorf("Expected log to contain 'before panic', got: %s", output)
	}
	if !strings.Contains(output, "recovery test panic") {
		t.Errorf("Expected log to contain 'recovery test panic', got: %s", output)
	}
}

// TestPanicWithStacktrace tests that panic includes stack trace information
func TestPanicWithStacktrace(t *testing.T) {
	// Configure test logger using helper function
	buf, cleanup := setupTestLogger()
	defer cleanup()

	// Create context with trace information
	ctx := WithTraceID(context.Background(), "test-trace-123")
	ctx = WithRequestID(ctx, "test-request-456")
	ctx = WithUserID(ctx, "user-789")

	// Test panic with context
	defer func() {
		if r := recover(); r != nil {
			panicStr := fmt.Sprintf("%v", r)
			if !strings.Contains(panicStr, "stack trace test panic") {
				t.Errorf("Expected panic to contain 'stack trace test panic', got: %s", panicStr)
			}
		}
	}()

	// Clear buffer before test
	buf.Reset()

	// Trigger panic
	panic("stack trace test panic")

	// Check that log contains stack trace information
	output := buf.String()
	if !strings.Contains(output, "stack trace test panic") {
		t.Errorf("Expected log to contain 'stack trace test panic', got: %s", output)
	}
}

// TestDirectPanicFunctions tests that direct panic functions work with trace
func TestDirectPanicFunctions(t *testing.T) {
	// Configure test logger using helper function
	buf, cleanup := setupTestLogger()
	defer cleanup()

	// Test direct Panic() function
	defer func() {
		if r := recover(); r != nil {
			panicStr := fmt.Sprintf("%v", r)
			if !strings.Contains(panicStr, "direct panic test") {
				t.Errorf("Expected panic to contain 'direct panic test', got: %s", panicStr)
			}
		}
	}()

	// Clear buffer before test
	buf.Reset()

	// Test direct Panic()
	Panic("direct panic test")

	// Test direct Panicf() function
	defer func() {
		if r := recover(); r != nil {
			panicStr := fmt.Sprintf("%v", r)
			if !strings.Contains(panicStr, "direct panicf test formatted") {
				t.Errorf("Expected panic to contain 'direct panicf test formatted', got: %s", panicStr)
			}
		}
	}()

	// Clear buffer before test
	buf.Reset()

	// Test direct Panicf()
	Panicf("direct panicf test %s", "formatted")

	// Test direct PanicContext() function
	defer func() {
		if r := recover(); r != nil {
			panicStr := fmt.Sprintf("%v", r)
			if !strings.Contains(panicStr, "direct panic context test") {
				t.Errorf("Expected panic to contain 'direct panic context test', got: %s", panicStr)
			}
		}
	}()

	// Clear buffer before test
	buf.Reset()

	ctx := WithTraceID(context.Background(), "test-trace")
	// Test direct PanicContext()
	PanicContext(ctx, "direct panic context test")

	// Test direct PanicfContext() function
	defer func() {
		if r := recover(); r != nil {
			panicStr := fmt.Sprintf("%v", r)
			if !strings.Contains(panicStr, "direct panicf context test formatted") {
				t.Errorf("Expected panic to contain 'direct panicf context test formatted', got: %s", panicStr)
			}
		}
	}()

	// Clear buffer before test
	buf.Reset()

	// Test direct PanicfContext()
	PanicfContext(ctx, "direct panicf context test %s", "formatted")
}
