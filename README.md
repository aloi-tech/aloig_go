# aloig - Modular Logging Library

`aloig` is a modular, extensible, and production-ready logging library based on [logrus](https://github.com/sirupsen/logrus) with [Sentry](https://sentry.io/) integration. Designed for Go applications that require a robust logging system with advanced capabilities.

## Features

- ✅ **Simple and structured interface** - Intuitive and easy-to-use API
- ✅ **Multi-environment support** - Different formats based on environment (JSON for production)
- ✅ **Sentry integration** - Automatic error reporting in production environments
- ✅ **Customizable fields** - Add contextual information to all your logs
- ✅ **Flexible configuration** - Adapt logger behavior to your needs
- ✅ **Singleton or multiple instances** - Use the global instance or create multiple loggers
- ✅ **Fully compatible with logrus** - Leverage the logrus ecosystem

## Installation

```bash
go get github.com/aloi-tech/aloig_go/aloig
```

## Basic Usage

### Default Logger (Singleton)

```go
import "github.com/aloi-tech/aloig_go/aloig"

func main() {
    // Get the singleton logger
    log := aloig.GetLogger()
    
    // Use the logger
    log.Info("Application started")
    log.WithField("user", "admin").Info("User connected")
}
```

### Custom Logger

```go
import (
    "github.com/aloi-tech/aloig_go/aloig"
    "github.com/sirupsen/logrus"
)

func main() {
    // Create custom configuration
    config := aloig.Config{
        Environment:      "staging",
        AppName:          "my-application",
        SentryDSN:        "https://your-dsn@sentry.io/123456",
        Level:            logrus.DebugLevel,
        ReportCaller:     true,
        CustomFields:     map[string]interface{}{"service": "api"},
    }
    
    // Create custom logger
    log := aloig.NewLogger(config)
    
    // Use the logger
    log.Debug("Configuration loaded")
}
```

## Configuration

The `Config` structure allows you to customize logger behavior:

```go
type Config struct {
    Environment      string                  // Environment: dev, staging, prod, etc.
    AppName          string                  // Application name
    SentryDSN        string                  // DSN for Sentry integration
    Release          string                  // Application version
    TracesSampleRate float64                 // Sampling rate for Sentry (0.0-1.0)
    Level            logrus.Level            // Minimum logging level
    ReportCaller     bool                    // Report the function that made the log
    CustomFields     map[string]interface{}  // Additional fields in all logs
}
```

### Default Configuration

The `DefaultConfig()` function creates a configuration based on environment variables:

- `ENVIRONMENT` - Application environment
- `APP_NAME` - Application name
- `SENTRY_DSN` - Sentry DSN
- `DEPLOY_ID` - Deployment ID (for versioning)

## Logging Levels

Available levels (from lowest to highest severity):

- `Debug` - Detailed information for debugging
- `Info` - General information about normal operation
- `Warn` - Warnings that don't interrupt operation
- `Error` - Errors that affect a specific operation
- `Fatal` - Critical errors that cause program termination
- `Panic` - Critical errors that cause a panic

## Advanced Features

### Context-Aware Logging

`aloig` provides context-aware logging functions that automatically include trace information:

```go
import (
    "context"
    "github.com/aloi-tech/aloig_go/aloig"
)

func handleRequest(ctx context.Context) {
    // Add trace information to context
    ctx = aloig.WithTraceID(ctx, "trace-123")
    ctx = aloig.WithRequestID(ctx, "req-456")
    ctx = aloig.WithUserID(ctx, "user-789")
    
    // Log with context - automatically includes trace information
    aloig.InfoContext(ctx, "Processing request")
    aloig.ErrorContext(ctx, "Database connection failed")
}
```

### Package-Level Functions

For convenience, `aloig` provides package-level functions that use the singleton logger:

```go
import "github.com/aloi-tech/aloig_go/aloig"

func main() {
    // Direct usage without getting logger instance
    aloig.Info("Application started")
    aloig.WithField("service", "api").Info("Service initialized")
    
    // Context-aware functions
    ctx := aloig.WithTraceID(context.Background(), "trace-123")
    aloig.InfoContext(ctx, "Request processed")
}
```

### Custom Fields and Chaining

```go
// Add custom fields
log := aloig.GetLogger()
log.WithField("user_id", "123").Info("User action")

// Chain multiple fields
log.WithFields(map[string]interface{}{
    "user_id": "123",
    "action":  "login",
    "ip":      "192.168.1.1",
}).Info("User logged in")

// Add error information
err := someFunction()
log.WithError(err).Error("Operation failed")
```

## Environment-Specific Behavior

### Development Environment

In development (`ENVIRONMENT=dev`):
- Uses text format for human-readable logs
- Includes colors and timestamps
- No Sentry integration

### Production Environment

In production (`ENVIRONMENT=prod`, `staging`, `sandbox`):
- Uses JSON format for structured logging
- Includes automatic fields (environment, app name, hostname, etc.)
- Sentry integration for error reporting
- Automatic stack traces for error levels

## Sentry Integration

When configured with a Sentry DSN, `aloig` automatically:
- Reports errors, fatal, and panic levels to Sentry
- Includes context information (trace ID, user ID, etc.)
- Provides stack traces for better debugging
- Flushes pending events on application shutdown

```go
// Configure with Sentry
config := aloig.Config{
    Environment: "prod",
    AppName:     "my-app",
    SentryDSN:   "https://your-dsn@sentry.io/123456",
}

log := aloig.NewLogger(config)

// Errors will be automatically reported to Sentry
log.Error("Database connection failed")
```

## Examples

See the `example/` directory for complete usage examples:

- `trace_example.go` - HTTP server with trace middleware
- `servicio_ejemplo.go` - Service example with custom configuration

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 