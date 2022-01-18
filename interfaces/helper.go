package interfaces

import (
	"encoding/json"
	"net/http"
)

func WriteHeader(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode("")
}
