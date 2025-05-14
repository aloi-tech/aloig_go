package aloig

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockLogger es una implementación simulada de la interfaz Logger para pruebas
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

// Métodos con contexto
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

// TestMockLogger demuestra cómo usar el mock para pruebas
func TestMockLogger(t *testing.T) {
	mockLog := new(MockLogger)

	// Configurar expectativas
	mockLog.On("Info", []interface{}{"test message"}).Return()
	mockLog.On("WithField", "key", "value").Return(mockLog)
	mockLog.On("Info", []interface{}{"with field"}).Return()

	// Usar el mock
	mockLog.Info("test message")
	mockLog.WithField("key", "value").Info("with field")

	// Verificar que se llamaron los métodos esperados
	mockLog.AssertExpectations(t)
}

// TestSentryIntegration simula la integración con Sentry
func TestSentryIntegration(t *testing.T) {
	// Este es un ejemplo de cómo podríamos probar la integración con Sentry
	// sin realmente conectarnos a Sentry, usando mocks y ajustes en el código

	// En una implementación real, podríamos necesitar exponer algunas funciones
	// o crear interfaces adicionales para facilitar el mockeo de Sentry

	// Ejemplo conceptual:
	mockLog := new(MockLogger)

	// Creamos un error real para la prueba
	testError := errors.New("error de prueba")

	// Configurar expectativas para un error que debería enviarse a Sentry
	mockLog.On("WithError", mock.Anything).Return(mockLog)
	mockLog.On("Error", []interface{}{"Error crítico"}).Return()

	// Simulamos un error que debería enviarse a Sentry
	mockLog.WithError(testError).Error("Error crítico")

	// Verificar expectativas
	mockLog.AssertExpectations(t)
}

// TestConfigureLogger verifica que ConfigureLogger funciona correctamente
func TestConfigureLogger(t *testing.T) {
	// Restablecer el singleton para la prueba
	log = nil
	once = sync.Once{}

	// Configurar con valores personalizados
	config := Config{
		Environment:  "test",
		AppName:      "mock-app",
		Level:        0, // Utiliza el nivel por defecto
		ReportCaller: true,
	}

	// Llamar a la función que estamos probando
	ConfigureLogger(config)

	// Obtener la instancia configurada
	logger := GetLogger()

	// Verificar que no es nil
	if logger == nil {
		t.Error("Logger configurado no debería ser nil")
	}

	// En una implementación real podríamos exponer alguna API para verificar
	// que los valores de configuración se aplicaron correctamente
}
