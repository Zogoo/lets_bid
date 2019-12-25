package v1

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": true, "message": "Hello world"}

	json.NewEncoder(w).Encode(resp)
}
