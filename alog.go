// Package alog proporciona una biblioteca de logging modular y extensible
// basada en logrus con integración a Sentry y capacidades avanzadas de logging.
//
// Esta biblioteca puede ser importada y utilizada en cualquier proyecto Go.
package alog

import (
	"os"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
	sentrylogrus "github.com/getsentry/sentry-go/logrus"
	"github.com/sirupsen/logrus"
)

// Logger es una interfaz que define las operaciones básicas de logging
// Esto permite reemplazar la implementación si es necesario
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	WithError(err error) Logger
}

// Config contiene la configuración para el logger
type Config struct {
	// Environment es el entorno actual (dev, staging, prod, etc.)
	Environment string

	// AppName es el nombre de la aplicación
	AppName string

	// SentryDSN es el DSN para la integración con Sentry
	SentryDSN string

	// Release es la versión de la aplicación
	Release string

	// TracesSampleRate es la tasa de muestreo para las trazas en Sentry (0.0 - 1.0)
	TracesSampleRate float64

	// Level es el nivel de logging mínimo
	Level logrus.Level

	// ReportCaller indica si se debe reportar la función que hizo el log
	ReportCaller bool

	// CustomFields son campos personalizados que se añadirán a todos los logs
	CustomFields map[string]interface{}
	HostName     string
}

// DefaultConfig crea una configuración por defecto
func DefaultConfig() Config {
	return Config{
		Environment:      os.Getenv("ENVIRONMENT"),
		AppName:          os.Getenv("APP_NAME"),
		SentryDSN:        os.Getenv("SENTRY_DSN"),
		Release:          os.Getenv("APP_NAME") + "@" + os.Getenv("DEPLOY_ID"),
		HostName:         os.Getenv("HOSTNAME"),
		TracesSampleRate: 0.2,
		Level:            logrus.InfoLevel,
		ReportCaller:     true,
		CustomFields:     make(map[string]interface{}),
	}
}

// FieldsHook es un hook para añadir campos personalizados a todos los logs
type FieldsHook struct {
	Fields logrus.Fields
}

// Levels devuelve los niveles a los que se aplicará el hook
func (hook *FieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire añade los campos personalizados a la entrada de log
func (hook *FieldsHook) Fire(entry *logrus.Entry) error {
	for key, value := range hook.Fields {
		entry.Data[key] = value
	}
	return nil
}

// logrusLogger es una implementación de Logger que usa logrus
type logrusLogger struct {
	logger *logrus.Logger
}

// isSentryEnvironment verifica si el entorno actual requiere integración con Sentry
func isSentryEnvironment(env string) bool {
	return env == "staging" || env == "sandbox" || env == "prod"
}

var (
	instance Logger
	once     sync.Once
)

// NewLogger crea una nueva instancia de Logger según la configuración proporcionada
func NewLogger(config Config) Logger {
	logrusInstance := logrus.New()

	// Configurar nivel de logging
	logrusInstance.SetLevel(config.Level)
	logrusInstance.SetReportCaller(config.ReportCaller)

	// Configurar formato según entorno
	if config.Environment != "dev" {
		logrusInstance.SetOutput(os.Stdout)
		standardFields := logrus.Fields{
			"env":      config.Environment,
			"appname":  config.AppName,
			"hostname": config.HostName,
		}

		// Añadir campos personalizados
		for k, v := range config.CustomFields {
			standardFields[k] = v
		}

		logrusInstance.AddHook(&FieldsHook{Fields: standardFields})
		logrusInstance.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrusInstance.SetOutput(os.Stdout)
		logrusInstance.SetFormatter(&logrus.TextFormatter{})
	}

	// Inicializar Sentry si es necesario
	if isSentryEnvironment(config.Environment) && config.SentryDSN != "" {
		err := initializeSentry(config)
		if err != nil {
			logrusInstance.WithError(err).Error("Error al inicializar Sentry")
		} else {
			// Configurar hook de Sentry
			sentryLevels := []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
			sentryHook, err := sentrylogrus.New(sentryLevels, sentry.CurrentHub().Client().Options())
			if err != nil {
				logrusInstance.WithError(err).Error("Error al crear hook de Sentry")
			} else {
				logrusInstance.AddHook(sentryHook)
				// Registrar manejador para flush de eventos en cierre
				logrus.RegisterExitHandler(func() {
					sentryHook.Flush(2 * time.Second)
				})
				logrusInstance.Info("Sentry inicializado correctamente")
			}
		}
	}

	return &logrusLogger{logger: logrusInstance}
}

// initializeSentry configura la conexión con Sentry
func initializeSentry(config Config) error {
	return sentry.Init(sentry.ClientOptions{
		Dsn:              config.SentryDSN,
		Environment:      config.Environment,
		Release:          config.Release,
		AttachStacktrace: true,
		TracesSampleRate: config.TracesSampleRate,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			return event
		},
	})
}

// GetLogger devuelve una instancia singleton del logger
func GetLogger() Logger {
	once.Do(func() {
		instance = NewLogger(DefaultConfig())
	})
	return instance
}

// ConfigureLogger configura la instancia singleton del logger con la configuración dada
func ConfigureLogger(config Config) {
	once.Do(func() {
		instance = NewLogger(config)
	})
}

// FlushSentry asegura que todos los eventos pendientes se envíen a Sentry
func FlushSentry() {
	if isSentryEnvironment(os.Getenv("ENVIRONMENT")) {
		sentry.Flush(2 * time.Second)
	}
}

// Implementación de la interfaz Logger para logrusLogger

func (l *logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *logrusLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *logrusLogger) WithField(key string, value interface{}) Logger {
	return &logrusLogger{logger: l.logger.WithField(key, value).Logger}
}

func (l *logrusLogger) WithFields(fields map[string]interface{}) Logger {
	logrusFields := logrus.Fields{}
	for k, v := range fields {
		logrusFields[k] = v
	}
	return &logrusLogger{logger: l.logger.WithFields(logrusFields).Logger}
}

func (l *logrusLogger) WithError(err error) Logger {
	return &logrusLogger{logger: l.logger.WithError(err).Logger}
}

// GetLogLevelFromEnv obtiene el nivel de log desde una variable de entorno
// Si la variable no existe o el valor no es válido, retorna el nivel por defecto
func GetLogLevelFromEnv(envVar, defaultLevel string) logrus.Level {
	levelStr := os.Getenv(envVar)
	if levelStr == "" {
		levelStr = defaultLevel
	}

	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		return logrus.InfoLevel
	}

	return level
}
