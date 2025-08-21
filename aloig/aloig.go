// Package aloig provides a modular and extensible logging library
// based on logrus with Sentry integration and advanced logging capabilities.
//
// This library can be imported and used in any Go project.
package aloig

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
	sentrylogrus "github.com/getsentry/sentry-go/logrus"
	"github.com/sirupsen/logrus"
)

// Logger is an interface that defines basic logging operations
// This allows replacing the implementation if necessary
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	WithError(err error) Logger
	WithContext(ctx context.Context) Logger

	// Context methods
	DebugContext(ctx context.Context, args ...interface{})
	DebugfContext(ctx context.Context, format string, args ...interface{})
	InfoContext(ctx context.Context, args ...interface{})
	InfofContext(ctx context.Context, format string, args ...interface{})
	WarnContext(ctx context.Context, args ...interface{})
	WarnfContext(ctx context.Context, format string, args ...interface{})
	WarningContext(ctx context.Context, args ...interface{})
	WarningfContext(ctx context.Context, format string, args ...interface{})
	ErrorContext(ctx context.Context, args ...interface{})
	ErrorfContext(ctx context.Context, format string, args ...interface{})
	FatalContext(ctx context.Context, args ...interface{})
	FatalfContext(ctx context.Context, format string, args ...interface{})
	PanicContext(ctx context.Context, args ...interface{})
	PanicfContext(ctx context.Context, format string, args ...interface{})
	PrintContext(ctx context.Context, args ...interface{})
	PrintfContext(ctx context.Context, format string, args ...interface{})
	PrintlnContext(ctx context.Context, args ...interface{})
	TraceContext(ctx context.Context, args ...interface{})
	TracefContext(ctx context.Context, format string, args ...interface{})
}

// Config contains the configuration for the logger
type Config struct {
	// Environment is the current environment (dev, staging, prod, etc.)
	Environment string

	// AppName is the application name
	AppName string

	// SentryDSN is the DSN for Sentry integration
	SentryDSN string

	// Release is the application version
	Release string

	// TracesSampleRate is the sampling rate for traces in Sentry (0.0 - 1.0)
	TracesSampleRate float64

	// Level is the minimum logging level
	Level logrus.Level

	// ReportCaller indicates whether to report the function that made the log
	ReportCaller bool

	// CustomFields are custom fields that will be added to all logs
	CustomFields map[string]interface{}
	HostName     string
	ServerName   string
}

// DefaultConfig creates a default configuration
func DefaultConfig() Config {
	return Config{
		Environment:      os.Getenv("ENVIRONMENT"),
		AppName:          os.Getenv("APP_NAME"),
		SentryDSN:        os.Getenv("SENTRY_DSN"),
		Release:          os.Getenv("APP_NAME") + "@" + os.Getenv("DEPLOY_ID"),
		HostName:         os.Getenv("HOSTNAME"),
		ServerName:       os.Getenv("APP_NAME"),
		TracesSampleRate: 0.2,
		Level:            logrus.TraceLevel,
		ReportCaller:     true,
		CustomFields:     make(map[string]interface{}),
	}
}

// FieldsHook is a hook to add custom fields to all logs
type FieldsHook struct {
	Fields logrus.Fields
}

// Levels returns the levels to which the hook will be applied
func (hook *FieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire adds custom fields to the log entry
func (hook *FieldsHook) Fire(entry *logrus.Entry) error {
	for key, value := range hook.Fields {
		entry.Data[key] = value
	}
	return nil
}

// CallerJSONFormatter is a custom JSON formatter that includes caller information
type CallerJSONFormatter struct {
	*logrus.JSONFormatter
}

// Format formats the log entry including caller information
func (f *CallerJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Get caller information
	if entry.Caller != nil {
		entry.Data["caller"] = fmt.Sprintf("%s:%d", filepath.Base(entry.Caller.File), entry.Caller.Line)
		entry.Data["function"] = getFunctionName(entry.Caller.Function)
		entry.Data["full_function"] = entry.Caller.Function
		entry.Data["file"] = entry.Caller.File
		entry.Data["line"] = entry.Caller.Line
	}

	// Add stack trace for error levels and above
	if entry.Level >= logrus.ErrorLevel {
		// Get stack trace
		stack := make([]byte, 4096)
		length := runtime.Stack(stack, false)
		stackStr := string(stack[:length])

		// Clean and format the stack trace
		lines := strings.Split(stackStr, "\n")
		var cleanStack []string
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" && !strings.Contains(line, "runtime/debug.Stack") &&
				!strings.Contains(line, "github.com/sirupsen/logrus") &&
				!strings.Contains(line, "aloig.(*CallerJSONFormatter).Format") {
				cleanStack = append(cleanStack, line)
			}
		}

		if len(cleanStack) > 0 {
			entry.Data["stack_trace"] = strings.Join(cleanStack, "\n")
		}
	}

	return f.JSONFormatter.Format(entry)
}

// getFunctionName extracts the function name without the package
func getFunctionName(fullName string) string {
	parts := strings.Split(fullName, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return fullName
}

// logrusLogger is a Logger implementation that uses logrus
type logrusLogger struct {
	logger *logrus.Logger
	ctx    context.Context
}

// isSentryEnvironment checks if the current environment requires Sentry integration
func isSentryEnvironment(env string) bool {
	return env == "staging" || env == "sandbox" || env == "prod"
}

var (
	log  Logger
	once sync.Once
)

// NewLogger creates a new Logger instance according to the provided configuration
func NewLogger(config Config) Logger {
	logrusInstance := logrus.New()

	// Configure logging level
	logrusInstance.SetLevel(config.Level)
	logrusInstance.SetReportCaller(config.ReportCaller)

	// Configure format according to environment
	if config.Environment != "dev" {
		logrusInstance.SetOutput(os.Stdout)
		standardFields := logrus.Fields{
			"env":        config.Environment,
			"appname":    config.AppName,
			"hostname":   config.HostName,
			"servername": config.ServerName,
			"release":    config.Release,
		}

		// Add custom fields
		for k, v := range config.CustomFields {
			standardFields[k] = v
		}

		logrusInstance.AddHook(&FieldsHook{Fields: standardFields})
		logrusInstance.SetFormatter(&CallerJSONFormatter{JSONFormatter: &logrus.JSONFormatter{}})
	} else {
		logrusInstance.SetOutput(os.Stdout)
		logrusInstance.SetFormatter(&logrus.TextFormatter{})
	}

	// Initialize Sentry if necessary
	if isSentryEnvironment(config.Environment) && config.SentryDSN != "" {
		err := initializeSentry(config)
		if err != nil {
			logrusInstance.WithError(err).Error("Error initializing Sentry")
		} else {
			// Configure Sentry hook
			sentryLevels := []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
			sentryHook, err := sentrylogrus.New(sentryLevels, sentry.CurrentHub().Client().Options())
			if err != nil {
				logrusInstance.WithError(err).Error("Error creating Sentry hook")
			} else {
				logrusInstance.AddHook(sentryHook)
				// Register handler for event flush on exit
				logrus.RegisterExitHandler(func() {
					sentryHook.Flush(2 * time.Second)
				})
				logrusInstance.Info("Sentry initialized successfully")
			}
		}
	}

	return &logrusLogger{logger: logrusInstance}
}

// initializeSentry configures the connection with Sentry
func initializeSentry(config Config) error {
	return sentry.Init(sentry.ClientOptions{
		Dsn:              config.SentryDSN,
		Environment:      config.Environment,
		Release:          config.Release,
		AttachStacktrace: true,
		ServerName:       config.AppName,
		TracesSampleRate: config.TracesSampleRate,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			return event
		},
		Tags: map[string]string{
			"env":        config.Environment,
			"appname":    config.AppName,
			"hostname":   config.HostName,
			"servername": config.ServerName,
			"release":    config.Release,
		},
	})
}

// GetLogger returns a singleton instance of the logger
func GetLogger() Logger {
	once.Do(func() {
		log = NewLogger(DefaultConfig())
	})
	return log
}

// ConfigureLogger configures the singleton logger instance with the given configuration
func ConfigureLogger(config Config) {
	once.Do(func() {
		log = NewLogger(config)
	})
}

// FlushSentry ensures that all pending events are sent to Sentry
func FlushSentry() {
	if isSentryEnvironment(os.Getenv("ENVIRONMENT")) {
		sentry.Flush(2 * time.Second)
	}
}

// Logger interface implementation for logrusLogger

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

func (l *logrusLogger) Warning(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Warningf(format string, args ...interface{}) {
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

func (l *logrusLogger) Print(args ...interface{}) {
	l.logger.Print(args...)
}

func (l *logrusLogger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *logrusLogger) Println(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *logrusLogger) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

func (l *logrusLogger) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

func (l *logrusLogger) WithField(key string, value interface{}) Logger {
	return &logrusLogger{logger: l.logger.WithField(key, value).Logger, ctx: l.ctx}
}

func (l *logrusLogger) WithFields(fields map[string]interface{}) Logger {
	logrusFields := logrus.Fields{}
	for k, v := range fields {
		logrusFields[k] = v
	}
	return &logrusLogger{logger: l.logger.WithFields(logrusFields).Logger, ctx: l.ctx}
}

func (l *logrusLogger) WithError(err error) Logger {
	return &logrusLogger{logger: l.logger.WithError(err).Logger, ctx: l.ctx}
}

func (l *logrusLogger) WithContext(ctx context.Context) Logger {
	return &logrusLogger{logger: l.logger, ctx: ctx}
}

// Context method implementation

func (l *logrusLogger) DebugContext(ctx context.Context, args ...interface{}) {
	l.withContextFields(ctx).Debug(args...)
}

func (l *logrusLogger) DebugfContext(ctx context.Context, format string, args ...interface{}) {
	l.withContextFields(ctx).Debugf(format, args...)
}

func (l *logrusLogger) InfoContext(ctx context.Context, args ...interface{}) {
	l.withContextFields(ctx).Info(args...)
}

func (l *logrusLogger) InfofContext(ctx context.Context, format string, args ...interface{}) {
	l.withContextFields(ctx).Infof(format, args...)
}

func (l *logrusLogger) WarnContext(ctx context.Context, args ...interface{}) {
	l.withContextFields(ctx).Warn(args...)
}

func (l *logrusLogger) WarnfContext(ctx context.Context, format string, args ...interface{}) {
	l.withContextFields(ctx).Warnf(format, args...)
}

func (l *logrusLogger) WarningContext(ctx context.Context, args ...interface{}) {
	l.withContextFields(ctx).Warning(args...)
}

func (l *logrusLogger) WarningfContext(ctx context.Context, format string, args ...interface{}) {
	l.withContextFields(ctx).Warningf(format, args...)
}

func (l *logrusLogger) ErrorContext(ctx context.Context, args ...interface{}) {
	l.withContextFields(ctx).Error(args...)
}

func (l *logrusLogger) ErrorfContext(ctx context.Context, format string, args ...interface{}) {
	l.withContextFields(ctx).Errorf(format, args...)
}

func (l *logrusLogger) FatalContext(ctx context.Context, args ...interface{}) {
	l.withContextFields(ctx).Fatal(args...)
}

func (l *logrusLogger) FatalfContext(ctx context.Context, format string, args ...interface{}) {
	l.withContextFields(ctx).Fatalf(format, args...)
}

func (l *logrusLogger) PanicContext(ctx context.Context, args ...interface{}) {
	l.withContextFields(ctx).Panic(args...)
}

func (l *logrusLogger) PanicfContext(ctx context.Context, format string, args ...interface{}) {
	l.withContextFields(ctx).Panicf(format, args...)
}

func (l *logrusLogger) PrintContext(ctx context.Context, args ...interface{}) {
	l.withContextFields(ctx).Print(args...)
}

func (l *logrusLogger) PrintfContext(ctx context.Context, format string, args ...interface{}) {
	l.withContextFields(ctx).Printf(format, args...)
}

func (l *logrusLogger) PrintlnContext(ctx context.Context, args ...interface{}) {
	l.withContextFields(ctx).Println(args...)
}

func (l *logrusLogger) TraceContext(ctx context.Context, args ...interface{}) {
	l.withContextFields(ctx).Trace(args...)
}

func (l *logrusLogger) TracefContext(ctx context.Context, format string, args ...interface{}) {
	l.withContextFields(ctx).Tracef(format, args...)
}

// withContextFields extracts context fields and adds them to the logger
func (l *logrusLogger) withContextFields(ctx context.Context) Logger {
	if ctx == nil {
		return l
	}

	fields := ExtractContextFields(ctx)
	if len(fields) == 0 {
		return l
	}

	return l.WithFields(fields)
}

// GetLogLevelFromEnv gets the log level from an environment variable
// If the variable doesn't exist or the value is invalid, returns the default level
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
