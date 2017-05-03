package golog

import (
	"log"
	"os"
)

var (
	std = &Logger{
		Logger:*log.New(os.Stdout, "", log.Ldate | log.Ltime | log.Lshortfile),
		level:TraceLevel,
		calldepth: 1}
)

func StandardLogger() *Logger {
	return std
}

func Level() int {
	return std.Level()
}

func SetLevel(level int) {
	std.SetLevel(level)
}

func SetLevelString(levelString string) {
	std.SetLevelString(levelString)
}

func Tracef(format string, v ...interface{}) {
	std.Tracef(format, v...)
}

func Trace(v ...interface{}) {
	std.Trace(v...)
}

func Traceln(v ...interface{}) {
	std.Traceln(v...)
}

func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

func Debug(v ...interface{}) {
	std.Debug(v...)
}

func Debugln(v ...interface{}) {
	std.Debugln(v...)
}

func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

func Info(v ...interface{}) {
	std.Info(v...)
}

func Infoln(v ...interface{}) {
	std.Infoln(v...)
}

func Warnf(format string, v ...interface{}) {
	std.Warnf(format, v...)
}

func Warn(v ...interface{}) {
	std.Warn(v...)
}

func Warnln(v ...interface{}) {
	std.Warnln(v...)
}

func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

func Error(v ...interface{}) {
	std.Error(v...)
}

func Errorln(v ...interface{}) {
	std.Errorln(v...)
}

func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}

func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

func Fatalln(v ...interface{}) {
	std.Fatalln(v...)
}

func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}

func Panic(v ...interface{}) {
	std.Panic(v...)
}

func Panicln(v ...interface{}) {
	std.Panicln(v...)
}

