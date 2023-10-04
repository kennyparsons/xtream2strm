package idsearch

import (
	"log"
	"strings"
	"xtream2strm/models"
	"xtream2strm/process"
)

func SearchVOD(query string, config models.Config) []SearchResult {
	vodData, err := process.GetVOD(config)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var results []SearchResult
	normalizedQuery := normalize(query)

	for _, vod := range vodData {
		normalizedTarget := normalize(vod.Name)

		// Check for direct substring match
		if strings.Contains(normalizedTarget, normalizedQuery) {
			distance := fuzzyMatch(normalizedQuery, normalizedTarget)
			results = append(results, SearchResult{
				ID:       vod.StreamID,
				Name:     vod.Name,
				Type:     "movie",
				Distance: distance,
			})
		}
	}

	return results
}

func SearchSeries(query string, config models.Config) []SearchResult {
	seriesData, err := process.GetSeries(config) // Assuming GetSeries returns SeriesJSON and error
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var results []SearchResult
	normalizedQuery := normalize(query)

	for _, series := range seriesData {
		normalizedTarget := normalize(series.Name)

		if strings.Contains(normalizedTarget, normalizedQuery) {
			distance := fuzzyMatch(normalizedQuery, normalizedTarget)
			results = append(results, SearchResult{
				ID:       series.SeriesID,
				Name:     series.Name,
				Type:     "series",
				Distance: distance,
			})
		}
	}
	return results
}
