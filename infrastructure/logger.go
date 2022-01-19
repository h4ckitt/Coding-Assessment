package infrastructure

import (
	"assessment/usecases"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Logger struct {
	Output *os.File
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
	log.SetOutput(io.MultiWriter(l.Output, os.Stderr))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format, v...)
}

func (l *Logger) LogAccess(r *http.Request, code int) {
	logFormat := fmt.Sprintf("[%s] - %s -> %d\n", r.Method, r.URL, code)

	log.SetOutput(io.MultiWriter(l.Output, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(logFormat)
}
