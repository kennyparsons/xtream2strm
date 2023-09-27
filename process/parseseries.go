package process

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"xtream2strm/models"
)

// CreateSeriesStrmFile creates an strm file for a given Series.
func CreateSeriesStrmFile(series models.Series, config models.Config) error {
	// Sanitize the name for the file name
	fileName := sanitizeFileName(series.Name) + ".strm"
	// Construct the file path
	filePath := filepath.Join(config.OutputDir, fileName)
	// Check if the output directory exists, if not, create it
	if _, err := os.Stat(config.OutputDir); os.IsNotExist(err) {
		os.MkdirAll(config.OutputDir, os.ModePerm)
	}

	// Construct the file content. This might be different for Series, adjust accordingly.
	fileContent := fmt.Sprintf("%s/series/%s/%s/%d",
		config.APIEndpoint, config.Username, config.Password, series.SeriesID)

	// Create and write to the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fileName, err)
	}
	defer file.Close()

	_, err = file.WriteString(fileContent)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", fileName, err)
	}

	return nil
}

// ParseSeriesData iterates through each Series and creates an strm file.
func ParseSeriesData(seriesData models.SeriesJSON, config models.Config) error {
	// Define a rate limit for the API calls of 1 per second
	rateLimit := time.NewTicker(1 * time.Second)
	defer rateLimit.Stop()

	for _, series := range seriesData {
		// Check if the Series's CategoryID is in the ignore list
		ignore := false
		for _, ignoreCategory := range config.IgnoreCategories {
			if series.CategoryID == ignoreCategory {
				ignore = true
				break
			}
		}
		if !ignore {
			// Wait for the rate limit tick
			<-rateLimit.C
			// Get the series info and iterate through the seasons and episodes
			// Log the series name
			log.Printf("Processing series: %s\n", series.Name)
			seriesInfo, err := GetSeriesInfo(fmt.Sprintf("%d", series.SeriesID), config)
			if err != nil {
				log.Printf("failed to get series info for %s: %v\n", series.Name, err)
				continue
			}
			// Show directory structure
			// create the folder for the season
			// the path is: <output_dir>/<series_name>/Season.<season_number>
			// Sanitize the series name for the folder name
			seriesFolderName := sanitizeFileName(series.Name)
			// Construct the folder path
			seriesFolderPath := filepath.Join(config.OutputDir, "tv", seriesFolderName)
			// Check if the series folder exists, if not, create it
			if _, err := os.Stat(seriesFolderPath); os.IsNotExist(err) {
				os.MkdirAll(seriesFolderPath, os.ModePerm)
			}

			for _, season := range seriesInfo.Seasons {
				// Construct the season folder path
				seasonFolderPath := filepath.Join(seriesFolderPath, fmt.Sprintf("Season.%d", season.SeasonNumber))
				// Check if the season folder exists, if not, create it
				if _, err := os.Stat(seasonFolderPath); os.IsNotExist(err) {
					os.MkdirAll(seasonFolderPath, os.ModePerm)
				}
				seasonID := fmt.Sprintf("%d", season.SeasonNumber)
				if episodes, ok := seriesInfo.Episodes[seasonID]; ok {
					for _, episode := range episodes {
						// Define a fileName for each episode
						fileName := fmt.Sprintf("S%02dE%02d.strm", season.SeasonNumber, episode.EpisodeNum)
						filePath := filepath.Join(seasonFolderPath, fileName)
						fileContent := fmt.Sprintf("%s/series/%s/%s/%s.%s", config.APIEndpoint, config.Username, config.Password, episode.ID, episode.ContainerExtension)
						// debug, print the file path and content to the screen
						// fmt.Printf("File Path: %s\n", filePath)
						// fmt.Printf("File Content: %s\n", fileContent)

						file, err := os.Create(filePath)
						if err != nil {
							return fmt.Errorf("failed to create file %s: %w", fileName, err)
						}

						_, err = file.WriteString(fileContent)
						if err != nil {
							file.Close()
							return fmt.Errorf("failed to write to file %s: %w", fileName, err)
						}
						file.Close()

						// // pause and ask the user to continue
						// fmt.Println("Press 'Enter' to continue...")
						// fmt.Scanln()
					}
				}
			}
		}
	}
	return nil
}

func GetSeriesInfo(seriesID string, config models.Config) (models.SeriesInfoResponse, error) {
	// Construct the URL
	url := fmt.Sprintf("%s/player_api.php?action=get_series_info&username=%s&password=%s&series_id=%s", config.APIEndpoint, config.Username, config.Password, seriesID)

	// Send the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return models.SeriesInfoResponse{}, fmt.Errorf("failed to fetch series info from Xtream Codes API: %w", err)
	}
	// Check the status code of the response to make sure it's 200 OK
	if resp.StatusCode != http.StatusOK {
		return models.SeriesInfoResponse{}, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var seriesInfo models.SeriesInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&seriesInfo)
	if err != nil {
		return models.SeriesInfoResponse{}, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	// // Enumerate through the seasons and episodes
	// for _, season := range seriesInfo.Seasons {
	// 	fmt.Printf("Season %d:\n", season.SeasonNumber)
	// 	seasonID := fmt.Sprintf("%d", season.SeasonNumber)
	// 	if episodes, ok := seriesInfo.Episodes[seasonID]; ok {
	// 		for _, episode := range episodes {
	// 			fmt.Printf("  Episode %d: %s (ID: %s)\n", episode.EpisodeNum, episode.Title, episode.ID)
	// 		}
	// 	}
	// }

	return seriesInfo, nil
}
