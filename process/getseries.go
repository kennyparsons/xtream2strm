package process

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"xtream2strm/models"
)

func GetSeries(config models.Config) (models.SeriesJSON, error) {
	// Construct the API URL for series
	apiURL := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=get_series", config.APIEndpoint, config.Username, config.Password)

	// Send a GET request to the Xtream Codes API
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch series data from Xtream Codes API: %w", err)
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request for series failed with status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read series response body: %w", err)
	}

	// Parse the JSON response
	var seriesData models.SeriesJSON
	err = json.Unmarshal(body, &seriesData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse series JSON response: %w", err)
	}

	// return the series data
	return seriesData, nil

	// // Process the series data (e.g., create strm files or perform other actions)
	// err = ParseSeriesData(seriesData, config)
	// if err != nil {
	// 	return fmt.Errorf("failed to parse series data: %w", err)
	// }

	// return nil
}
