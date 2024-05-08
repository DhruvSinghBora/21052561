// internal/server/server.go
package server

import (
	"average-calculator/internal/numbers"
	"encoding/json"
	// "fmt"
	"net/http"
)

func Start() {
	http.HandleFunc("/numbers/", numbersHandler)
	http.ListenAndServe(":9876", nil)
}

func numbersHandler(w http.ResponseWriter, r *http.Request) {
	// Extract number ID from URL path
	numberID := r.URL.Path[len("/numbers/"):]

	// Fetch numbers from the test server
	numbers, err := numbers.Fetch(numberID)
	if err != nil {
		http.Error(w, "Failed to fetch numbers", http.StatusInternalServerError)
		return
	}

	// Calculate average of the fetched numbers
	avg := calculateAverage(numbers)

	// Construct response object
	response := struct {
		Numbers         []int     `json:"numbers"`
		WindowPrevState []int     `json:"windowPrevState"`
		WindowCurrState []int     `json:"windowCurrState"`
		Avg             float64   `json:"avg"`
	}{
		Numbers:         numbers,
		WindowPrevState: []int{},
		WindowCurrState: numbers,
		Avg:             avg,
	}

	// Marshal response object to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	// Set content type header and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func calculateAverage(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}
