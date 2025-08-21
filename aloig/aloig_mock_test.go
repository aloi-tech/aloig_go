package aloig

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockLogger is a mock implementation of the Logger interface for testing
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(args ...interface{}) {
	m.Called(args)
}

func (m *MockLogger) Debugf(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) Info(args ...interface{}) {
	m.Called(args)
}

func (m *MockLogger) Infof(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) Warn(args ...interface{}) {
	m.Called(args)
}

func (m *MockLogger) Warning(args ...interface{}) {
	m.Called(args)
}

func (m *MockLogger) Warnf(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) Warningf(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) Error(args ...interface{}) {
	m.Called(args)
}

func (m *MockLogger) Errorf(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) Fatal(args ...interface{}) {
	m.Called(args)
}

func (m *MockLogger) Fatalf(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) Panic(args ...interface{}) {
	m.Called(args)
}

func (m *MockLogger) Panicf(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) Printf(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) Print(args ...interface{}) {
	m.Called(args)
}

func (m *MockLogger) Println(args ...interface{}) {
	m.Called(args)
}

func (m *MockLogger) Trace(args ...interface{}) {
	m.Called(args)
}

func (m *MockLogger) Tracef(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockLogger) WithField(key string, value interface{}) Logger {
	args := m.Called(key, value)
	return args.Get(0).(Logger)
}

func (m *MockLogger) WithFields(fields map[string]interface{}) Logger {
	args := m.Called(fields)
	return args.Get(0).(Logger)
}

func (m *MockLogger) WithError(err error) Logger {
	args := m.Called(err)
	return args.Get(0).(Logger)
}

func (m *MockLogger) WithContext(ctx context.Context) Logger {
	args := m.Called(ctx)
	return args.Get(0).(Logger)
}

// Context methods
func (m *MockLogger) DebugContext(ctx context.Context, args ...interface{}) {
	m.Called(ctx, args)
}

func (m *MockLogger) DebugfContext(ctx context.Context, format string, args ...interface{}) {
	m.Called(ctx, format, args)
}

func (m *MockLogger) InfoContext(ctx context.Context, args ...interface{}) {
	m.Called(ctx, args)
}

func (m *MockLogger) InfofContext(ctx context.Context, format string, args ...interface{}) {
	m.Called(ctx, format, args)
}

func (m *MockLogger) WarnContext(ctx context.Context, args ...interface{}) {
	m.Called(ctx, args)
}

func (m *MockLogger) WarnfContext(ctx context.Context, format string, args ...interface{}) {
	m.Called(ctx, format, args)
}

func (m *MockLogger) WarningContext(ctx context.Context, args ...interface{}) {
	m.Called(ctx, args)
}

func (m *MockLogger) WarningfContext(ctx context.Context, format string, args ...interface{}) {
	m.Called(ctx, format, args)
}

func (m *MockLogger) ErrorContext(ctx context.Context, args ...interface{}) {
	m.Called(ctx, args)
}

func (m *MockLogger) ErrorfContext(ctx context.Context, format string, args ...interface{}) {
	m.Called(ctx, format, args)
}

func (m *MockLogger) FatalContext(ctx context.Context, args ...interface{}) {
	m.Called(ctx, args)
}

func (m *MockLogger) FatalfContext(ctx context.Context, format string, args ...interface{}) {
	m.Called(ctx, format, args)
}

func (m *MockLogger) PanicContext(ctx context.Context, args ...interface{}) {
	m.Called(ctx, args)
}

func (m *MockLogger) PanicfContext(ctx context.Context, format string, args ...interface{}) {
	m.Called(ctx, format, args)
}

func (m *MockLogger) PrintContext(ctx context.Context, args ...interface{}) {
	m.Called(ctx, args)
}

func (m *MockLogger) PrintfContext(ctx context.Context, format string, args ...interface{}) {
	m.Called(ctx, format, args)
}

func (m *MockLogger) PrintlnContext(ctx context.Context, args ...interface{}) {
	m.Called(ctx, args)
}

func (m *MockLogger) TraceContext(ctx context.Context, args ...interface{}) {
	m.Called(ctx, args)
}

func (m *MockLogger) TracefContext(ctx context.Context, format string, args ...interface{}) {
	m.Called(ctx, format, args)
}

// TestMockLoggerBasicFunctions tests basic mock logger functions
func TestMockLoggerBasicFunctions(t *testing.T) {
	mockLogger := &MockLogger{}

	// Set up expectations
	mockLogger.On("Info", mock.AnythingOfType("[]interface {}")).Return()
	mockLogger.On("Error", mock.AnythingOfType("[]interface {}")).Return()
	mockLogger.On("Debug", mock.AnythingOfType("[]interface {}")).Return()

	// Call functions
	mockLogger.Info("test message")
	mockLogger.Error("test error")
	mockLogger.Debug("test debug")

	// Verify expectations
	mockLogger.AssertExpectations(t)
}

// TestMockLoggerWithFields tests mock logger with fields
func TestMockLoggerWithFields(t *testing.T) {
	mockLogger := &MockLogger{}
	mockReturnLogger := &MockLogger{}

	// Set up expectations
	mockLogger.On("WithField", "key", "value").Return(mockReturnLogger)
	mockReturnLogger.On("Info", mock.AnythingOfType("[]interface {}")).Return()

	// Call functions
	logger := mockLogger.WithField("key", "value")
	logger.Info("test with field")

	// Verify expectations
	mockLogger.AssertExpectations(t)
	mockReturnLogger.AssertExpectations(t)
}

// TestMockLoggerContextFunctions tests mock logger context functions
func TestMockLoggerContextFunctions(t *testing.T) {
	mockLogger := &MockLogger{}
	ctx := context.Background()

	// Set up expectations
	mockLogger.On("InfoContext", ctx, mock.AnythingOfType("[]interface {}")).Return()
	mockLogger.On("ErrorContext", ctx, mock.AnythingOfType("[]interface {}")).Return()

	// Call functions
	mockLogger.InfoContext(ctx, "test context message")
	mockLogger.ErrorContext(ctx, "test context error")

	// Verify expectations
	mockLogger.AssertExpectations(t)
}

// TestMockLoggerChaining tests mock logger chaining
func TestMockLoggerChaining(t *testing.T) {
	mockLogger := &MockLogger{}
	mockReturnLogger1 := &MockLogger{}
	mockReturnLogger2 := &MockLogger{}

	// Set up expectations for chaining
	mockLogger.On("WithField", "field1", "value1").Return(mockReturnLogger1)
	mockReturnLogger1.On("WithError", mock.AnythingOfType("*errors.errorString")).Return(mockReturnLogger2)
	mockReturnLogger2.On("Info", mock.AnythingOfType("[]interface {}")).Return()

	// Call chained functions
	testError := errors.New("test error")
	mockLogger.WithField("field1", "value1").
		WithError(testError).
		Info("chained message")

	// Verify expectations
	mockLogger.AssertExpectations(t)
	mockReturnLogger1.AssertExpectations(t)
	mockReturnLogger2.AssertExpectations(t)
}

// TestMockLoggerSingletonReplacement tests replacing the singleton logger with a mock
func TestMockLoggerSingletonReplacement(t *testing.T) {
	// Save original logger
	originalLog := log

	// Create mock logger
	mockLogger := &MockLogger{}
	mockLogger.On("Info", mock.AnythingOfType("[]interface {}")).Return()

	// Replace singleton
	log = mockLogger

	// Call package-level function
	Info("test singleton replacement")

	// Verify expectations
	mockLogger.AssertExpectations(t)

	// Restore original logger
	log = originalLog
}

// TestMockLoggerConcurrentAccess tests mock logger with concurrent access
func TestMockLoggerConcurrentAccess(t *testing.T) {
	mockLogger := &MockLogger{}

	// Set up expectations for concurrent calls
	mockLogger.On("Info", mock.AnythingOfType("[]interface {}")).Return().Times(3)

	// Run concurrent calls
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mockLogger.Info("concurrent message", id)
		}(i)
	}

	wg.Wait()

	// Verify expectations
	mockLogger.AssertExpectations(t)
}
