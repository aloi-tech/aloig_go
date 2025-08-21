package example

import (
	"errors"
	"os"
	"time"

	aloig "github.com/aloi-tech/aloig_go/aloig"
)

// ExampleService demonstrates how to use alog in a real service
type ExampleService struct {
	logger aloig.Logger
}

// NewExampleService creates a new service instance with its own configured logger
func NewExampleService() *ExampleService {
	// Custom configuration for this service
	config := aloig.Config{
		Environment:  os.Getenv("ENVIRONMENT"),
		AppName:      "example-service",
		SentryDSN:    os.Getenv("SENTRY_DSN"),
		Level:        aloig.GetLogLevelFromEnv("LOG_LEVEL", "info"),
		HostName:     os.Getenv("HOSTNAME"),
		ReportCaller: true,
		CustomFields: map[string]interface{}{
			"module":  "example-service",
			"version": "1.0.0",
		},
	}

	return &ExampleService{
		logger: aloig.NewLogger(config),
	}
}

// Start starts the service and logs information
func (s *ExampleService) Start() error {
	s.logger.Info("Starting example service")

	// Example of using additional fields
	s.logger.WithField("timestamp", time.Now().Unix()).Info("Service initialized successfully")

	return nil
}

// Process simulates a process with possible errors to demonstrate error logging
func (s *ExampleService) Process(data string) error {
	logger := s.logger.WithFields(map[string]interface{}{
		"operation": "process",
		"data":      data,
	})

	logger.Debug("Starting data processing")

	// Simulation of an error
	if data == "" {
		err := errors.New("empty data")
		logger.WithError(err).Error("Error processing empty data")
		return err
	}

	logger.Info("Data processed successfully")
	return nil
}

// Finish finalizes the service and ensures all logs are sent
func (s *ExampleService) Finish() {
	s.logger.Info("Finishing service")
	// Ensure all Sentry messages are sent before exiting
	aloig.FlushSentry()
}
