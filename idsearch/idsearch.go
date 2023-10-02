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
	// order the results by distance. Highest distance first.
	return results
}

// func SearchVOD(query string, vodData models.XtreamCodesJSON) []SearchResult {
//     var results []SearchResult

//     for _, vod := range vodData {
//         distance := fuzzyMatch(query, vod.Name)
//         if distance <= len(query)/2 { // or any other threshold you find suitable
//             results = append(results, SearchResult{
//                 ID:   vod.StreamID,
//                 Name: vod.Name,
//                 Type: "movie",
//             })
//         }
//     }

//     return results
// }

func SearchSeries(query string, seriesData models.SeriesJSON) []SearchResult {
	var results []SearchResult

	for _, series := range seriesData {
		distance := fuzzyMatch(query, series.Name)
		if distance <= len(query)/2 { // or any other threshold you find suitable
			results = append(results, SearchResult{
				ID:   series.SeriesID,
				Name: series.Name,
				Type: "series",
			})
		}
	}

	return results
}
