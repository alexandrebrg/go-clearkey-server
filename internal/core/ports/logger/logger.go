package logger

var Log logger

type logger interface {
	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})

	Infof(format string, args ...interface{})
	Info(args ...interface{})

	Warnf(format string, args ...interface{})

	Debugf(format string, args ...interface{})
	Debug(args ...interface{})

	Printf(format string, args ...interface{})
}

func SetLogger(newLogger logger) {
	Log = newLogger
}
