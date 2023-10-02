package process

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"xtream2strm/models"
	// other necessary imports
)

func GetVOD(config models.Config) (models.XtreamCodesJSON, error) {
	// Construct the API URL using the APIEndpoint from the config
	apiURL := fmt.Sprintf("%s/player_api.php?username=%s&password=%s&action=get_vod_streams", config.APIEndpoint, config.Username, config.Password)

	// Send a GET request to the Xtream Codes API
	resp, err := http.Get(apiURL)
	if err != nil {
		return models.XtreamCodesJSON{}, fmt.Errorf("failed to fetch data from Xtream Codes API: %w", err)
	}
	// Check the status code of the response to make sure it's 200 OK
	if resp.StatusCode != http.StatusOK {
		return models.XtreamCodesJSON{}, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.XtreamCodesJSON{}, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var xtreamData models.XtreamCodesJSON
	err = json.Unmarshal(body, &xtreamData)
	if err != nil {
		return models.XtreamCodesJSON{}, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return xtreamData, nil

	// // Process the VOD data and create Jellyfin compatible strm files
	// err = ParseVODData(xtreamData, config)
	// if err != nil {
	// 	return fmt.Errorf("failed to parse VOD data: %w", err)
	// }

	// return nil
}
