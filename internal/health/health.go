package health

import (
	"encoding/json"
	"net/http"
	"time"
)


type Response struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`   
}

func CheckHealth(version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := Response {
			Status:    "OK",
			Timestamp: time.Now(),
			Version:   version,
		}
		 
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}