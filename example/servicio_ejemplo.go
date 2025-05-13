package ejemplo

import (
	"errors"
	"os"
	"time"

	aloig "github.com/aloi/alog"
)

// ServicioEjemplo demuestra cómo usar alog en un servicio real
type ServicioEjemplo struct {
	logger aloig.Logger
}

// NuevoServicioEjemplo crea una nueva instancia del servicio con su propio logger configurado
func NuevoServicioEjemplo() *ServicioEjemplo {
	// Configuración personalizada para este servicio
	config := aloig.Config{
		Environment:  os.Getenv("ENVIRONMENT"),
		AppName:      "servicio-ejemplo",
		SentryDSN:    os.Getenv("SENTRY_DSN"),
		Level:        aloig.GetLogLevelFromEnv("LOG_LEVEL", "info"),
		HostName:     os.Getenv("HOSTNAME"),
		ReportCaller: true,
		CustomFields: map[string]interface{}{
			"modulo":  "servicio-ejemplo",
			"version": "1.0.0",
		},
	}

	return &ServicioEjemplo{
		logger: aloig.NewLogger(config),
	}
}

// Iniciar inicia el servicio y registra información en el log
func (s *ServicioEjemplo) Iniciar() error {
	s.logger.Info("Iniciando servicio de ejemplo")

	// Ejemplo de uso de campos adicionales
	s.logger.WithField("timestamp", time.Now().Unix()).Info("Servicio inicializado correctamente")

	return nil
}

// Procesar simula un proceso con posibles errores para demostrar el logging de errores
func (s *ServicioEjemplo) Procesar(datos string) error {
	logger := s.logger.WithFields(map[string]interface{}{
		"operacion": "procesar",
		"datos":     datos,
	})

	logger.Debug("Iniciando procesamiento de datos")

	// Simulación de un error
	if datos == "" {
		err := errors.New("datos vacíos")
		logger.WithError(err).Error("Error al procesar datos vacíos")
		return err
	}

	logger.Info("Datos procesados correctamente")
	return nil
}

// Finalizar finaliza el servicio y asegura que todos los logs se envíen
func (s *ServicioEjemplo) Finalizar() {
	s.logger.Info("Finalizando servicio")
	// Aseguramos que todos los mensajes de Sentry se envíen antes de salir
	aloig.FlushSentry()
}
