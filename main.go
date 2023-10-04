package main

import (
	"flag"
	"fmt"
	"log"

	"xtream2strm/config"
	"xtream2strm/idsearch"
	"xtream2strm/process"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to the config file")
	searchTerm := flag.String("search", "", "search for a movie or series")
	restrictTo := flag.String("restrict-to", "", "restrict operations to specific types (e.g., movies,tv)")
	flag.Parse()

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	process.ParseRestrictTo(*restrictTo)

	// --search will be followed by a search term
	if *searchTerm != "" {
		fmt.Println("Searching for", *searchTerm)
		// check if this operation is allowed for movies
		if process.IsOperationAllowed("movies") {
			fmt.Println("Searching for", *searchTerm, "in movies...")
			vodresults := idsearch.SearchVOD(*searchTerm, config)
			idsearch.DisplaySearchResults(vodresults)
		}
		if process.IsOperationAllowed("tv") {
			fmt.Println("Searching for", *searchTerm, "in series...")
			seriesresults := idsearch.SearchSeries(*searchTerm, config)
			idsearch.DisplaySearchResults(seriesresults)
		}

		// quit the program after displaying the search results
		return

	}
	if process.IsOperationAllowed("movies") {
		xtreamData, err := process.GetVOD(config)
		if err != nil {
			log.Fatal(err)
		}
		err = process.ParseVODData(xtreamData, config)
		if err != nil {
			log.Fatal(err)
		}
	}
	if process.IsOperationAllowed("tv") {
		seriesData, err := process.GetSeries(config)
		if err != nil {
			log.Fatal(err)
		}
		err = process.ParseSeriesData(seriesData, config)
		if err != nil {
			log.Fatal(err)
		}
	}
	// // Register the FileHandler to handle incoming requests
	// http.HandleFunc("/", server.FileHandler)

	// // Add the root directory to the virtual file system
	// server.AddToFileSystem("/", models.VirtualFile{IsDir: true})
	// server.AddToFileSystem("/tv/", models.VirtualFile{IsDir: true})

	// // Start the HTTP server
	// fmt.Print("Starting server on :8089...\n")
	// go func() {
	// 	err := http.ListenAndServe(":8089", nil)
	// 	if err != nil {
	// 		log.Fatal("Server Failure: %v", err)
	// 	}
	// }()

	// err = process.GetVOD(config)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // err = process.GetSeries(config)
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

	// fmt.Println("All movies and shows have been processed successfully.")
	// //Prevent the program from exiting to allow the server to continue serving requests
	// select {}
}
