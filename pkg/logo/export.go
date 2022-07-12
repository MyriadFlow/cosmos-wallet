package logo

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

//Logs with error field
func ErrorWField(err error, args ...interface{}) {
	log.WithField("msg", err.Error()).Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

//Logs with formatting and error field
func ErrorfWField(format string, err error, args ...interface{}) {
	log.WithField("msg", err.Error()).Errorf(format, args...)
}
