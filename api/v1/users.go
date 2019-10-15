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
		return
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
		return
	}

	userID, err := service.AuthenticateWithPassword(user)

	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid credentials"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	token := service.GenerateNewToken(userID)

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = token

	json.NewEncoder(w).Encode(resp)
}
