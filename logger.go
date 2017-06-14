package golog

import (
	"log"
	"fmt"
	"os"
	"sync"
	"strings"
)

const (
	TraceLevel = iota
	DebugLevel
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

var (
	LevelString = []string{"[TRACE] ", "[DEBUG] ", "[INFO] ", "[WARN] ", "[ERROR] ", "[FATAL] ", "[PANIC] "}
)

type Logger struct {
	log.Logger
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	LogLevel int
	Calldepth int
}

func (l *Logger) Level() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.LogLevel
}

func (l *Logger) SetLevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.LogLevel = level
}

func (l *Logger) SetLevelString(levelString string) {
	switch strings.ToLower(levelString) {
	case "trace":
		l.SetLevel(TraceLevel)
	case "debug":
		l.SetLevel(DebugLevel)
	case "info":
		l.SetLevel(InfoLevel)
	case "warn":
		l.SetLevel(WarningLevel)
	case "error":
		l.SetLevel(ErrorLevel)
	case "fatal":
		l.SetLevel(FatalLevel)
	case "panic":
		l.SetLevel(PanicLevel)
	}
}

func (l *Logger) Logf(level int, format string, v ...interface{}) {
	if l.LogLevel > level {
		return
	}

	s := fmt.Sprintf(format, v...)
	l.Output(3 + l.Calldepth, LevelString[level] + s)

	if level == PanicLevel {
		panic(s)
	} else if level == FatalLevel {
		os.Exit(1)
	}
}

func (l *Logger) Log(level int, v ...interface{}) {
	if l.LogLevel > level {
		return
	}

	s := fmt.Sprint(v...)
	l.Output(3 + l.Calldepth, LevelString[level] + s)

	if level == PanicLevel {
		panic(s)
	} else if level == FatalLevel {
		os.Exit(1)
	}
}

func (l *Logger) Logln(level int, v ...interface{}) {
	if l.LogLevel > level {
		return
	}

	s := fmt.Sprintln(v...)
	l.Output(3 + l.Calldepth, LevelString[level] + s)

	if level == PanicLevel {
		panic(s)
	} else if level == FatalLevel {
		os.Exit(1)
	}
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.Logf(TraceLevel, format, v...)
}

func (l *Logger) Trace(v ...interface{}) {
	l.Log(TraceLevel, v...)
}

func (l *Logger) Traceln(v ...interface{}) {
	l.Logln(TraceLevel, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logf(DebugLevel, format, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.Log(DebugLevel, v...)
}

func (l *Logger) Debugln(v ...interface{}) {
	l.Logln(DebugLevel, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logf(InfoLevel, format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.Log(InfoLevel, v...)
}

func (l *Logger) Infoln(v ...interface{}) {
	l.Logln(InfoLevel, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Logf(WarningLevel, format, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.Log(WarningLevel, v...)
}

func (l *Logger) Warnln(v ...interface{}) {
	l.Logln(WarningLevel, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logf(ErrorLevel, format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.Log(ErrorLevel, v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	l.Logln(ErrorLevel, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Logf(FatalLevel, format, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Log(FatalLevel, v...)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.Logln(FatalLevel, v...)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Logf(PanicLevel, format, v...)
}

func (l *Logger) Panic(v ...interface{}) {
	l.Log(PanicLevel, v...)
}

func (l *Logger) Panicln(v ...interface{}) {
	l.Logln(PanicLevel, v...)
}
