package process

import (
	"fmt"
	"xtream2strm/models"
	"xtream2strm/server"
)

func vfsMovie(vod models.VODStream, config models.Config) error {
	// Extract the necessary information from vod and config
	baseName := sanitizeFileName(vod.Name)
	fileName := baseName + "." + vod.ContainerExtension
	filePath := "/movies/" + baseName + "/" + fileName
	directLink := fmt.Sprintf("%s/movie/%s/%s/%d.%s",
		config.APIEndpoint, config.Username, config.Password, vod.StreamID, vod.ContainerExtension)

	// Add the parent directories to the virtualFS map as directories
	server.AddToFileSystem("/movies/", models.VirtualFile{IsDir: true})
	server.AddToFileSystem("/movies/"+baseName+"/", models.VirtualFile{IsDir: true})

	// Add the extracted information to the virtualFS map in the server package
	server.AddToFileSystem(filePath, models.VirtualFile{IsDir: false, DirectLink: directLink})

	return nil
}
