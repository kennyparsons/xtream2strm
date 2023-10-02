package idsearch

type SearchResult struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"` // "movie" or "series"
}