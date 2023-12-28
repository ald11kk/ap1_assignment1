package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestData struct {
	Message string `json:"message"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if requestData.Message != "" {
		fmt.Println("Received Message:", requestData.Message)

		response := ResponseData{Status: "success", Message: "Data successfully received"}
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response := ResponseData{Status: "400", Message: "Invalid JSON message"}
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	http.HandleFunc("/", handlePostRequest)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
