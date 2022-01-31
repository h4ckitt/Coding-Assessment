package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Logger struct {
	ErrorOutput  io.Writer
	AccessOutput io.Writer
}

func NewLogger() *Logger {
	return &Logger{os.Stderr, os.Stdout}
}

func NewTestLogger() *Logger {
	var buff bytes.Buffer
	return &Logger{&buff, &buff}
}

func (l *Logger) SetOutput(filepath string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		l.LogError("%s\n", err)
		return
	}

	l.ErrorOutput = file
}

func (l *Logger) LogError(format string, v ...interface{}) {
	//log.SetOutput(io.MultiWriter(l.ErrorOutput, os.Stderr))
	log.SetOutput(l.ErrorOutput)
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format, v...)
}

func (l *Logger) LogAccess(r *http.Request, code int) {
	logFormat := fmt.Sprintf("[%s] - %s -> %d\n", r.Method, r.URL, code)

	//log.SetOutput(io.MultiWriter(l.ErrorOutput, os.Stdout))
	log.SetOutput(l.AccessOutput)
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(logFormat)
}
