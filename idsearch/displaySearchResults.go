package idsearch

import (
	"fmt"
)

func DisplaySearchResults(results []SearchResult) {
	if len(results) == 0 {
		fmt.Println("No results found.")
		return
	}

	fmt.Println("---------------------------------------------------")
	fmt.Println("| ID       | Name                                               | Type  | Distance |")
	fmt.Println("---------------------------------------------------")
	for _, result := range results {
		fmt.Printf("| %-7d | %-50s | %-5s | %-8d |\n", result.ID, result.Name, result.Type, result.Distance)
	}
	fmt.Println("---------------------------------------------------")
}
