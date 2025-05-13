# aloig - Biblioteca de Logging Modular

`aloig` es una biblioteca de logging modular, extensible y lista para producción basada en [logrus](https://github.com/sirupsen/logrus) con integración de [Sentry](https://sentry.io/). Diseñada para ser utilizada en aplicaciones Go que requieren un sistema de logging robusto con capacidades avanzadas.

## Características

- ✅ **Interfaz simple y estructurada** - API intuitiva y fácil de usar
- ✅ **Soporte para entornos múltiples** - Formatos diferentes según el entorno (JSON para producción)
- ✅ **Integración con Sentry** - Reporte automático de errores en entornos productivos
- ✅ **Campos personalizables** - Añade información contextual a todos tus logs
- ✅ **Configuración flexible** - Adapta el comportamiento del logger a tus necesidades
- ✅ **Singleton o múltiples instancias** - Usa la instancia global o crea múltiples loggers
- ✅ **Totalmente compatible con logrus** - Aprovecha el ecosistema de logrus

## Instalación

```bash
go get github.com/aloi-tech/aloig_go/aloig
```

## Uso Básico

### Logger por defecto (Singleton)

```go
import "github.com/aloi-tech/aloig_go/aloig"

func main() {
    // Obtener el logger singleton
    log := aloig.GetLogger()
    
    // Usar el logger
    log.Info("Aplicación iniciada")
    log.WithField("usuario", "admin").Info("Usuario conectado")
}
```

### Logger personalizado

```go
import (
    "github.com/aloi-tech/aloig_go/aloig"
    "github.com/sirupsen/logrus"
)

func main() {
    // Crear configuración personalizada
    config := aloig.Config{
        Environment:      "staging",
        AppName:          "mi-aplicacion",
        SentryDSN:        "https://tu-dsn@sentry.io/123456",
        Level:            logrus.DebugLevel,
        ReportCaller:     true,
        CustomFields:     map[string]interface{}{"servicio": "api"},
    }
    
    // Crear logger personalizado
    log := aloig.NewLogger(config)
    
    // Usar el logger
    log.Debug("Configuración cargada")
}
```

## Configuración

La estructura `Config` permite personalizar el comportamiento del logger:

```go
type Config struct {
    Environment      string                  // Entorno: dev, staging, prod, etc.
    AppName          string                  // Nombre de la aplicación
    SentryDSN        string                  // DSN para la integración con Sentry
    Release          string                  // Versión de la aplicación
    TracesSampleRate float64                 // Tasa de muestreo para Sentry (0.0-1.0)
    Level            logrus.Level            // Nivel mínimo de logging
    ReportCaller     bool                    // Reportar la función que hizo el log
    CustomFields     map[string]interface{}  // Campos adicionales en todos los logs
}
```

### Configuración por defecto

La función `DefaultConfig()` crea una configuración basada en variables de entorno:

- `ENVIRONMENT` - Entorno de la aplicación
- `APP_NAME` - Nombre de la aplicación
- `SENTRY_DSN` - DSN para Sentry
- `DEPLOY_ID` - ID del despliegue (para versionado)

## Niveles de Logging

Los niveles disponibles son (de menor a mayor severidad):

- `Debug` - Información detallada para depuración
- `Info` - Información general sobre el funcionamiento normal
- `Warn` - Advertencias que no interrumpen el funcionamiento
- `Error` - Errores que afectan a una operación específica
- `Fatal` - Errores graves que provocan la terminación del programa
- `Panic` - Errores críticos que provocan un panic

## Integración con Sentry

La biblioteca integra automáticamente con Sentry en entornos de producción (`staging`, `sandbox` o `prod`). Los eventos de nivel `Error`, `Fatal` y `Panic` se envían a Sentry.

Para asegurar que todos los eventos se envíen antes de que termine el programa, usa:

```go
defer aloig.FlushSentry()
```

## Ejemplo Completo

```go
package main

import (
    "errors"
    "github.com/aloi-tech/aloig_go/aloig"
)

func main() {
    // Obtener el logger
    log := aloig.GetLogger()
    
    // Logging básico
    log.Info("Aplicación iniciada")
    
    // Logging con campos adicionales
    log.WithFields(map[string]interface{}{
        "usuario_id": 12345,
        "accion":     "login",
    }).Info("Usuario ha iniciado sesión")
    
    // Logging de errores
    err := errors.New("error de conexión")
    log.WithError(err).Error("No se pudo conectar a la base de datos")
    
    // Asegurar que los eventos de Sentry se envíen
    defer aloig.FlushSentry()
}
```

## Personalización Avanzada

### Agregar Hooks de logrus

Si necesitas extender la funcionalidad con hooks personalizados:

```go
import (
    "github.com/aloi-tech/aloig_go/aloig"
    "github.com/sirupsen/logrus"
)

// Tu hook personalizado
type CustomHook struct{}

func (h *CustomHook) Levels() []logrus.Level {
    return logrus.AllLevels
}

func (h *CustomHook) Fire(entry *logrus.Entry) error {
    // Lógica personalizada
    return nil
}

// Obtener la instancia subyacente de logrus
func getLogrusInstance(logger aloig.Logger) *logrus.Logger {
    // Nota: Esto requiere conocimiento interno de la implementación
    logrusLogger, ok := logger.(*aloig.logrusLogger)
    if !ok {
        return nil
    }
    return logrusLogger.logger
}
```

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un issue o envía un pull request.

## Licencia

[MIT](LICENSE) 