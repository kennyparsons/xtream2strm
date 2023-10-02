package idsearch

import (
    "github.com/agnivade/levenshtein"
)

func fuzzyMatch(query string, target string) int {
    return levenshtein.ComputeDistance(query, target)
}