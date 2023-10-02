package main

import (
	"flag"
	"log"

	"xtream2strm/config"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to the config file")
	flag.Parse()

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
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
