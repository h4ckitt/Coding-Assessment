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

//NewLogger : This Constructor Returns A New Instance Of The Logger Entity
func NewLogger() *Logger {
	return &Logger{os.Stderr, os.Stdout}
}

/*NewTestLogger : Returns A New Instance Of The Logger Entity
Whose Outputs Have Been Set To A Buffer To Prevent Messy Outputs
During Tests
*/
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

/*LogError : Writes Out Provided Errors To The Predefined Location
params:
	- format <string>
	- v <interface{}>
*/
func (l *Logger) LogError(format string, v ...interface{}) {
	//log.SetOutput(io.MultiWriter(l.ErrorOutput, os.Stderr))
	log.SetOutput(l.ErrorOutput)
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(format, v...)
}

/*LogAccess : Writes Out Details About Requests That Hit An Endpoint
params:
	- *http.Request
	- code <int>
*/
func (l *Logger) LogAccess(r *http.Request, code int) {
	logFormat := fmt.Sprintf("[%s] - %s -> %d\n", r.Method, r.URL, code)

	//log.SetOutput(io.MultiWriter(l.ErrorOutput, os.Stdout))
	log.SetOutput(l.AccessOutput)
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf(logFormat)
}
