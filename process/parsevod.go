package process

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"xtream2strm/models"
)

// sanitizeFileName replaces spaces with periods and removes all special characters from a string.
func sanitizeFileName(name string) string {
	// Replace spaces with periods
	name = strings.ReplaceAll(name, " ", ".")

	// Remove all special characters using a regular expression
	reg, err := regexp.Compile("[^a-zA-Z0-9.]+")
	if err != nil {
		return name // Return the original name if the regex compilation fails
	}

	name = reg.ReplaceAllString(name, "")

	// if name ends in a period, remove it
	if strings.HasSuffix(name, ".") {
		name = name[:len(name)-1]
	}

	// if the name starts with a period, remove it
	if strings.HasPrefix(name, ".") {
		name = name[1:]
	}

	// if there are any repeat periods, replace them with a single period.
	for strings.Contains(name, "..") {
		name = strings.ReplaceAll(name, "..", ".")
	}

	return name
}

// CreateStrmFile creates an strm file for a given VOD stream.
func CreateStrmFile(vod models.VODStream, config models.Config) error {
	// Sanitize the name for the file name
	fileName := sanitizeFileName(vod.Name) + ".strm"
	// Construct the file path
	filePath := filepath.Join(config.OutputDir, "movies", fileName)
	// Check if the output directory exists, if not, create it
	if _, err := os.Stat(config.OutputDir); os.IsNotExist(err) {
		os.MkdirAll(config.OutputDir, os.ModePerm)
	}
	// Check if the movies directory exists, if not, create it
	if _, err := os.Stat(filepath.Join(config.OutputDir, "movies")); os.IsNotExist(err) {
		os.MkdirAll(filepath.Join(config.OutputDir, "movies"), os.ModePerm)
	}

	// Construct the file content
	fileContent := fmt.Sprintf("%s/movie/%s/%s/%d.%s",
		config.APIEndpoint, config.Username, config.Password, vod.StreamID, vod.ContainerExtension)

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

// ParseVODData iterates through each VOD stream and creates an strm file.
func ParseVODData(vodData models.XtreamCodesJSON, config models.Config) error {
	for _, vod := range vodData {
		// Check if the VODStream's CategoryID is in the ignore list
		ignore := false
		for _, ignoreCategory := range config.IgnoreCategories {
			if vod.CategoryID == ignoreCategory {
				ignore = true
				break
			}
		}
		if !ignore {
			err := CreateStrmFile(vod, config)
			if err != nil {
				// log the error for this stream and continue
				log.Printf("Failed to create strm file for stream %s: %v\n", vod.Name, err)
			}
		}
	}
	return nil
}
