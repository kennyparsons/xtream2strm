package idsearch

import (
	"strings"

	"github.com/agnivade/levenshtein"
)

func fuzzyMatch(query string, target string) int {
	//debug
	// fmt.Println("query:", query)
	// fmt.Println("target:", target)
	fuzzydistance := levenshtein.ComputeDistance(query, target)
	// fmt.Println("fuzzydistance:", fuzzydistance)

	return fuzzydistance
}

func normalize(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}
