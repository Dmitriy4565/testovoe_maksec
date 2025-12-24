package main

import (
	"encoding/json"
	"net/http"
	"time"

	"testovoe_maksec/internal/checker"
)

type CheckRequest struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST", http.StatusMethodNotAllowed)
		return
	}

	var req CheckRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	res := checker.CheckPort(r.Context(), req.IP, req.Port, 3*time.Second)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/check", handler)
	http.ListenAndServe(":8080", nil)
}
