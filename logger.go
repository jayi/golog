package golog

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
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
	mu        sync.Mutex // ensures atomic writes; protects the following fields
	Level     int
	Calldepth int
}

func NewLogger(out io.Writer, level int) *Logger {
	return &Logger{
		Logger:    *log.New(out, "", log.Ldate|log.Ltime|log.Lshortfile),
		Level:     level,
		Calldepth: 0}
}

func NewFileLogger(filename string, level int, flag int) (*Logger, error) {
	fw, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	logger := &Logger{
		Logger:    *log.New(fw, "", flag),
		Level:     level,
		Calldepth: 0,
	}
	runtime.SetFinalizer(logger, func(logger *Logger) {
		fw.Close()
	})
	return logger, err
}

func (l *Logger) GetLevel() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Level
}

func (l *Logger) SetLevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Level = level
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
	if l.Level > level {
		return
	}

	s := fmt.Sprintf(format, v...)
	l.Output(3+l.Calldepth, LevelString[level]+s)

	if level == PanicLevel {
		panic(s)
	} else if level == FatalLevel {
		os.Exit(1)
	}
}

func (l *Logger) Log(level int, v ...interface{}) {
	if l.Level > level {
		return
	}

	s := fmt.Sprint(v...)
	l.Output(3+l.Calldepth, LevelString[level]+s)

	if level == PanicLevel {
		panic(s)
	} else if level == FatalLevel {
		os.Exit(1)
	}
}

func (l *Logger) Logln(level int, v ...interface{}) {
	if l.Level > level {
		return
	}

	s := fmt.Sprintln(v...)
	l.Output(3+l.Calldepth, LevelString[level]+s)

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

func (l *Logger) PrintJson(v ...interface{}) {
	for _, i := range v {
		data, err := json.Marshal(i)
		if err != nil {
			l.Println(i, err)
		} else {
			l.Println(string(data))
		}
	}
}
