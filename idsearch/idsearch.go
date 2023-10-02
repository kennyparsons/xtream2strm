package idsearch

import (
    "xtream2strm/models"
)

func SearchVOD(query string, vodData models.XtreamCodesJSON) []SearchResult {
    var results []SearchResult

    for _, vod := range vodData {
        distance := fuzzyMatch(query, vod.Name)
        if distance <= len(query)/2 { // or any other threshold you find suitable
            results = append(results, SearchResult{
                ID:   vod.StreamID,
                Name: vod.Name,
                Type: "movie",
            })
        }
    }

    return results
}

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