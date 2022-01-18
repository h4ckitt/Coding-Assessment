package infrastructure

import (
	"assessment/usecases"
	"io"
	"log"
	"os"
)

type Logger struct {
	Output os.File
}

func NewLogger() usecases.Logger {
	return &Logger{os.Stdout}
}

func (l *Logger) SetOutput(filepath string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		l.LogError("%s\n", err)
		return
	}

	l.Output = file
}

func (l *Logger) LogError(format string, v ...interface{}) {
	defer l.Output.Close()

	log.SetOutput(io.MultiWriter(l.Output, os.Stderr))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format, v...)
}
