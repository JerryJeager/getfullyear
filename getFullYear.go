// Package getFullYear provides functionality to retrieve the current year
// and related details from the GetFullYear API.
package getFullYear

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// getFullYear represents the structure of the response from the GetFullYear API.
type getFullYear struct {
	Year        int32  `json:"year"`         // Year represents the current year as an integer.
	SponsoredBy string `json:"sponsored_by"` // SponsoredBy indicates the sponsor of the year data.
	YearString  string `json:"year_string"`  // YearString is a human-readable representation of the year.
}

// endpoint defines the URL for the GetFullYear API.
var endpoint = "https://getfullyear.com/api/year"

// GetFullYear retrieves the current year details from the GetFullYear API.
//
// It makes an HTTP GET request to the API, parses the JSON response, and
// returns the year details as a getFullYear struct. If an error occurs
// during the HTTP request or JSON parsing, it returns an error.
//
// Returns:
//   - *getFullYear: A pointer to the getFullYear struct containing year details.
//   - error: An error if the HTTP request or JSON parsing fails.
//
// Example usage:
//   yearData, err := GetFullYear()
//   if err != nil {
//       log.Fatalf("Error fetching year data: %v", err)
//   }
//   fmt.Printf("Year: %d, Sponsored By: %s\n", yearData.Year, yearData.SponsoredBy)
func GetFullYear() (*getFullYear, error) {
	var fullYear getFullYear

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Create a new GET request.
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Set headers for the request.
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "GetFullYear/1.0")

	// Perform the HTTP request.
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check for non-200 status codes.
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return nil, errors.New("failed to get full year")
	}

	// Parse the JSON response.
	if err := json.Unmarshal(body, &fullYear); err != nil {
		return nil, err
	}

	return &fullYear, nil
}
