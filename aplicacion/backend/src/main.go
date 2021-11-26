package main

import (
    "net/http"
)

type HandsOn struct{
	Time time.time `json:"time"`
	Hostname string `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	resp = HandsOn{
		Time:time.now(),
		Hostname:os.Getenv("HOSTNAME"),
	}
	jsonResp, err := json.Marshal(&resp)

	w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write()
}

func main() {
    http.HandleFunc("/", ServeHTTP)
    http.ListenAndServe(":9090", nil)
}

