package main

import (
	"context"
	"net/http"
	"os"

	aloig "github.com/aloi-tech/aloig_go/aloig"
)

// Middleware para asignar un trace ID a cada solicitud
func traceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtener trace ID desde el header, o generar uno nuevo si no existe
		traceID := r.Header.Get("X-Trace-ID")
		if traceID == "" {
			traceID = aloig.GenerateTraceID()
			// Agregar el trace ID a la respuesta para propósitos de depuración
			w.Header().Set("X-Trace-ID", traceID)
		}

		// Crear un contexto con el trace ID
		ctx := aloig.WithTraceID(r.Context(), traceID)

		// Agregar ID de usuario si está disponible (por ejemplo, de una cookie o token)
		if userID := r.Header.Get("X-User-ID"); userID != "" {
			ctx = aloig.WithUserID(ctx, userID)
		}

		// Asignar un ID de request único
		requestID := aloig.GenerateTraceID()
		ctx = aloig.WithRequestID(ctx, requestID)

		// Continuar con el siguiente handler usando el contexto enriquecido
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Handler de ejemplo que usa el contexto para logging
func homeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Log usando el contexto, automáticamente incluirá trace_id, user_id, etc.
	aloig.InfoContext(ctx, "Petición recibida a la ruta principal")

	// Llamar a un servicio de negocio
	response, err := businessService(ctx)
	if err != nil {
		aloig.ErrorContext(ctx, "Error en el servicio de negocio")
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	aloig.InfoContext(ctx, "Petición procesada correctamente")
	w.Write([]byte(response))
}

// Simulación de un servicio de negocio que procesa la solicitud
func businessService(ctx context.Context) (string, error) {
	// Los logs aquí también incluirán automáticamente el trace ID
	aloig.InfoContext(ctx, "Procesando lógica de negocio")

	// Simulamos una llamada a un servicio externo
	result, err := externalServiceCall(ctx)
	if err != nil {
		aloig.ErrorContext(ctx, "Error al llamar al servicio externo")
		return "", err
	}

	return result, nil
}

// Simulación de una llamada a un servicio externo
func externalServiceCall(ctx context.Context) (string, error) {
	// El trace ID sigue acompañando a través de toda la cadena de llamadas
	aloig.InfoContext(ctx, "Llamando a servicio externo")

	// Devolvemos una respuesta simulada
	return "Respuesta del servicio externo", nil
}

func main() {
	// Configurar el entorno (opcional)
	os.Setenv("ENVIRONMENT", "dev")
	os.Setenv("APP_NAME", "ejemplo-trace")

	// Configurar rutas con el middleware de trace
	mux := http.NewServeMux()
	mux.Handle("/", traceMiddleware(http.HandlerFunc(homeHandler)))

	// Iniciar servidor
	aloig.Info("Iniciando servidor en :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		aloig.Fatal("Error al iniciar servidor:", err)
	}
}
