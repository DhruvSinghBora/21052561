// internal/numbers/numbers.go
package numbers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	testServerBaseURL = "http://20.244.56.144/test/"
	authToken         = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiZXhwIjoxNzE1MTQ1MjQzLCJpYXQiOjE3MTUxNDQ5NDMsImlzcyI6IkFmZm9yZG1lZCIsImp0aSI6ImMzY2YzZDM0LTlhYWEtNDg5NS1hM2YwLWU2ZDdiMDY1Njk0ZiIsInN1YiI6IjIxMDUyNzA1QGtpaXQuYWMuaW4ifSwiY29tcGFueU5hbWUiOiJBd2FzdGhpIiwiY2xpZW50SUQiOiJjM2NmM2QzNC05YWFhLTQ4OTUtYTNmMC1lNmQ3YjA2NTY5NGYiLCJjbGllbnRTZWNyZXQiOiJ3SWlaeWVHRkFxWlhDenZGIiwib3duZXJOYW1lIjoic2hpdmFtIiwib3duZXJFbWFpbCI6IjIxMDUyNzA1QGtpaXQuYWMuaW4iLCJyb2xsTm8iOiIyMTA1MjcwNSJ9.1tk2EZLxsDufmTZ6sicS6CZajIZ05LpYG6o0ODI-Ozs"
)

type TestServerResponse struct {
	Numbers []int `json:"numbers"`
}

func Fetch(numberID string) ([]int, error) {
	// Prepare the request URL
	requestURL := getRequestURL(numberID)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// Send HTTP request
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	// Add authorization header
	req.Header.Set("Authorization", authToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse response
	var testServerResponse TestServerResponse
	err = json.NewDecoder(resp.Body).Decode(&testServerResponse)
	if err != nil {
		return nil, err
	}

	return testServerResponse.Numbers, nil
}

func getRequestURL(numberID string) string {
	var apiURL string
	switch numberID {
	case "p":
		apiURL = testServerBaseURL + "primes"
	case "f":
		apiURL = testServerBaseURL + "fibo"
	case "e":
		apiURL = testServerBaseURL + "even"
	case "r":
		apiURL = testServerBaseURL + "random"
	default:
		panic(fmt.Sprintf("Unknown number ID: %s", numberID))
	}
	return apiURL
}
