package usecases

import "net/http"

type Logger interface {
	LogError(string, ...interface{})
	LogAccess(*http.Request, int)
}
