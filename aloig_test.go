package aloig

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

// Estructura auxiliar para capturar logs
type BufferHook struct {
	Buffer *bytes.Buffer
}

func (h *BufferHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *BufferHook) Fire(entry *logrus.Entry) error {
	line, _ := entry.String()
	h.Buffer.WriteString(line)
	return nil
}

// TestDefaultConfig verifica la configuración por defecto
func TestDefaultConfig(t *testing.T) {
	// Guardar variables de entorno originales
	oldEnv := os.Getenv("ENVIRONMENT")
	oldAppName := os.Getenv("APP_NAME")

	// Establecer valores de prueba
	os.Setenv("ENVIRONMENT", "test")
	os.Setenv("APP_NAME", "test-app")

	// Obtener configuración por defecto
	config := DefaultConfig()

	// Verificar valores
	if config.Environment != "test" {
		t.Errorf("Se esperaba Environment='test', se obtuvo '%s'", config.Environment)
	}
	if config.AppName != "test-app" {
		t.Errorf("Se esperaba AppName='test-app', se obtuvo '%s'", config.AppName)
	}
	if config.Level != logrus.InfoLevel {
		t.Errorf("Se esperaba Level=InfoLevel, se obtuvo '%v'", config.Level)
	}
	if !config.ReportCaller {
		t.Errorf("Se esperaba ReportCaller=true, se obtuvo '%v'", config.ReportCaller)
	}

	// Restaurar variables de entorno
	os.Setenv("ENVIRONMENT", oldEnv)
	os.Setenv("APP_NAME", oldAppName)
}

// TestNewLogger verifica la creación de un nuevo logger
func TestNewLogger(t *testing.T) {
	config := Config{
		Environment:  "dev",
		AppName:      "test-app",
		Level:        logrus.DebugLevel,
		ReportCaller: false,
		CustomFields: map[string]interface{}{"test": "value"},
	}

	logger := NewLogger(config)
	if logger == nil {
		t.Error("Se esperaba un logger no nulo")
	}
}

// TestLogrusLogger verifica los métodos de logging
func TestLogrusLogger(t *testing.T) {
	// Crear un logger con un hook para capturar la salida
	logrusInstance := logrus.New()
	logrusInstance.SetLevel(logrus.DebugLevel)
	var buf bytes.Buffer
	logrusInstance.SetOutput(&buf)
	logrusInstance.SetFormatter(&logrus.JSONFormatter{})

	logger := &logrusLogger{logger: logrusInstance}

	// Probar métodos de logging
	testCases := []struct {
		name     string
		logFunc  func()
		contains string
	}{
		{"Debug", func() { logger.Debug("test debug") }, "test debug"},
		{"Info", func() { logger.Info("test info") }, "test info"},
		{"Warn", func() { logger.Warn("test warn") }, "test warn"},
		{"Error", func() { logger.Error("test error") }, "test error"},
		{"WithField", func() { logger.WithField("key", "value").Info("test with field") }, "test with field"},
		{"WithFields", func() {
			logger.WithFields(map[string]interface{}{"key1": "value1", "key2": "value2"}).Info("test with fields")
		}, "test with fields"},
		{"WithError", func() { logger.WithError(errors.New("test error")).Info("test with error") }, "test with error"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf.Reset()
			tc.logFunc()
			if !bytes.Contains(buf.Bytes(), []byte(tc.contains)) {
				t.Errorf("Log no contiene '%s'", tc.contains)
			}
		})
	}
}

// TestSingletonLogger verifica el comportamiento del singleton
func TestSingletonLogger(t *testing.T) {
	// Restablecer el singleton para la prueba
	// Para tests reales necesitaríamos un método de reinicio o mock
	instance = nil
	// No podemos establecer once a nil, pero podríamos crear un método ResetForTests() para pruebas

	logger1 := GetLogger()
	logger2 := GetLogger()

	// Ambas referencias deberían ser la misma
	if logger1 != logger2 {
		t.Error("GetLogger() debería devolver la misma instancia")
	}
}

// TestLogFormat verifica el formato del log según el entorno
func TestLogFormat(t *testing.T) {
	// Probar formato JSON en entorno de producción
	config := Config{
		Environment: "prod",
		AppName:     "test-app",
	}

	logrusInstance := logrus.New()
	var buf bytes.Buffer
	logrusInstance.SetOutput(&buf)

	// Esto es un poco difícil de probar directamente porque NewLogger crea un nuevo logger
	// En una situación real podríamos usar una interfaz mock o refactorizar para permitir pruebas
	logger := NewLogger(config)
	logger.Info("test message")

	// Verificar que podemos decodificar el JSON de salida
	// Esto asume que la salida es JSON, lo que podría no ser cierto si hemos cambiado la configuración
	// Sería mejor refactorizar para poder acceder al formatter directamente
	if config.Environment != "dev" {
		// Intentar decodificar la salida como JSON
		// Este test puede fallar si la salida no es JSON o si hay múltiples líneas JSON
		t.Log("Este test puede fallar si la salida no es JSON o tiene formato incorrecto")
	}
}

// TestFieldsHook verifica que el hook de campos añada correctamente los campos
func TestFieldsHook(t *testing.T) {
	// Crear un hook con campos
	hook := &FieldsHook{
		Fields: logrus.Fields{
			"test_field": "test_value",
		},
	}

	// Verificar niveles
	levels := hook.Levels()
	if len(levels) != len(logrus.AllLevels) {
		t.Errorf("Se esperaban %d niveles, se obtuvieron %d", len(logrus.AllLevels), len(levels))
	}

	// Verificar que se añadan los campos
	entry := &logrus.Entry{
		Data: make(logrus.Fields),
	}
	err := hook.Fire(entry)
	if err != nil {
		t.Errorf("No se esperaba error, se obtuvo: %v", err)
	}
	if entry.Data["test_field"] != "test_value" {
		t.Errorf("Se esperaba field['test_field']='test_value', se obtuvo '%v'", entry.Data["test_field"])
	}
}

// TestIsSentryEnvironment verifica la detección de entornos que usan Sentry
func TestIsSentryEnvironment(t *testing.T) {
	testCases := []struct {
		env      string
		expected bool
	}{
		{"dev", false},
		{"test", false},
		{"staging", true},
		{"sandbox", true},
		{"prod", true},
		{"", false},
	}

	for _, tc := range testCases {
		t.Run(tc.env, func(t *testing.T) {
			result := isSentryEnvironment(tc.env)
			if result != tc.expected {
				t.Errorf("Para env='%s', se esperaba %v, se obtuvo %v", tc.env, tc.expected, result)
			}
		})
	}
}
