package aloig

// Este archivo contiene funciones de conveniencia a nivel de paquete
// para permitir el uso de aloig.Info(), aloig.Error(), etc. directamente

// Debug registra un mensaje de nivel debug usando el logger singleton
func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

// Debugf registra un mensaje de nivel debug con formato usando el logger singleton
func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

// Info registra un mensaje de nivel info usando el logger singleton
func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

// Infof registra un mensaje de nivel info con formato usando el logger singleton
func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

// Warn registra un mensaje de nivel warning usando el logger singleton
func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

// Warnf registra un mensaje de nivel warning con formato usando el logger singleton
func Warnf(format string, args ...interface{}) {
	GetLogger().Warnf(format, args...)
}

// Error registra un mensaje de nivel error usando el logger singleton
func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

// Errorf registra un mensaje de nivel error con formato usando el logger singleton
func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

// Fatal registra un mensaje de nivel fatal usando el logger singleton
// y luego hace que la aplicación termine con un status code no-cero
func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

// Fatalf registra un mensaje de nivel fatal con formato usando el logger singleton
// y luego hace que la aplicación termine con un status code no-cero
func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

// Panic registra un mensaje de nivel panic usando el logger singleton
// y luego lanza un panic con el mensaje formateado
func Panic(args ...interface{}) {
	GetLogger().Panic(args...)
}

// Panicf registra un mensaje de nivel panic con formato usando el logger singleton
// y luego lanza un panic con el mensaje formateado
func Panicf(format string, args ...interface{}) {
	GetLogger().Panicf(format, args...)
}

// WithField retorna una nueva entrada de log con el campo key=value añadido
func WithField(key string, value interface{}) Logger {
	return GetLogger().WithField(key, value)
}

// WithFields retorna una nueva entrada de log con los campos añadidos
func WithFields(fields map[string]interface{}) Logger {
	return GetLogger().WithFields(fields)
}

// WithError retorna una nueva entrada de log con un error añadido
func WithError(err error) Logger {
	return GetLogger().WithError(err)
}
