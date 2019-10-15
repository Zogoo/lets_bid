package v1

import (
	"encoding/json"
	"lets_bid/service"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	user := &service.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
	}
	service.CreateNewUser(user)

	json.NewEncoder(w).Encode("OK")
}

func Auth(w http.ResponseWriter, r *http.Request) {
	user := &service.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
	}
	service.AuthenticatePass(user)
	json.NewEncoder(w).Encode("OK")
}
