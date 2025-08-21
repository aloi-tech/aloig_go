package example

import (
	"context"
	"net/http"
	"os"

	aloig "github.com/aloi-tech/aloig_go/aloig"
)

// Middleware to assign a trace ID to each request
func traceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get trace ID from header, or generate a new one if it doesn't exist
		traceID := r.Header.Get("X-Trace-ID")
		if traceID == "" {
			traceID = aloig.GenerateTraceID()
			// Add trace ID to response for debugging purposes
			w.Header().Set("X-Trace-ID", traceID)
		}

		// Create a context with the trace ID
		ctx := aloig.WithTraceID(r.Context(), traceID)

		// Add user ID if available (e.g., from a cookie or token)
		if userID := r.Header.Get("X-User-ID"); userID != "" {
			ctx = aloig.WithUserID(ctx, userID)
		}

		// Assign a unique request ID
		requestID := aloig.GenerateTraceID()
		ctx = aloig.WithRequestID(ctx, requestID)

		// Continue with the next handler using the enriched context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Example handler that uses context for logging
func homeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Log using context, automatically includes trace_id, user_id, etc.
	aloig.InfoContext(ctx, "Request received to main route")

	// Call a business service
	response, err := businessService(ctx)
	if err != nil {
		aloig.ErrorContext(ctx, "Error in business service")
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// Successful response
	aloig.InfoContext(ctx, "Request processed successfully")
	w.Write([]byte(response))
}

// Simulation of a business service that processes the request
func businessService(ctx context.Context) (string, error) {
	// Logs here will also automatically include the trace ID
	aloig.InfoContext(ctx, "Processing business logic")

	// Simulate a call to an external service
	result, err := externalServiceCall(ctx)
	if err != nil {
		aloig.ErrorContext(ctx, "Error calling external service")
		return "", err
	}

	return result, nil
}

// Simulation of a call to an external service
func externalServiceCall(ctx context.Context) (string, error) {
	// Trace ID continues to accompany through the entire call chain
	aloig.InfoContext(ctx, "Calling external service")

	// Return a simulated response
	return "External service response", nil
}

func main() {
	// Configure environment (optional)
	os.Setenv("ENVIRONMENT", "dev")
	os.Setenv("APP_NAME", "trace-example")

	// Configure routes with trace middleware
	mux := http.NewServeMux()
	mux.Handle("/", traceMiddleware(http.HandlerFunc(homeHandler)))

	// Start server
	aloig.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		aloig.Fatal("Error starting server:", err)
	}
}
