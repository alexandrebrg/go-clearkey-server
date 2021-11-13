package loggers

import "go.uber.org/zap"


type loggerWrapper struct {
	lw *zap.SugaredLogger
}

func (logger *loggerWrapper) Errorf(format string, args ...interface{}) {
	logger.lw.Errorf(format, args...)
}

func (logger *loggerWrapper) Fatalf(format string, args ...interface{}) {
	logger.lw.Fatalf(format, args...)
}

func (logger *loggerWrapper) Fatal(args ...interface{}) {
	logger.lw.Fatal(args...)
}

func (logger *loggerWrapper) Infof(format string, args ...interface{}) {
	logger.lw.Infof(format, args...)
}

func (logger *loggerWrapper) Info(args ...interface{}) {
	logger.lw.Info(args...)
}


func (logger *loggerWrapper) Warnf(format string, args ...interface{}) {
	logger.lw.Warnf(format, args...)
}

func (logger *loggerWrapper) Debugf(format string, args ...interface{}) {
	logger.lw.Debugf(format, args...)
}

func (logger *loggerWrapper) Debug(args ...interface{}) {
	logger.lw.Debug(args...)
}

func (logger *loggerWrapper) Printf(format string, args ...interface{}) {
	logger.lw.Infof(format, args...)
}

func (logger *loggerWrapper) Println(args ...interface{}) {
	logger.lw.Info(args, "\n")
}

func NewZLogger(envType string) *loggerWrapper {
	var zlogger *zap.Logger

	switch envType {
	case "production":
		zlogger, _ = zap.NewProduction()
		break
	default:
		zlogger, _ = zap.NewDevelopment()
		break
	}

	defer zlogger.Sync()

	sugar := zlogger.Sugar()
	return &loggerWrapper{lw: sugar}
}
