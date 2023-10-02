package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"xtream2strm/models"
)

// Define the directory where files will be cached
const cacheDir = "/tmp/cache"

// virtualFS is a concurrent map to represent the virtual file system.
var virtualFS = struct {
	sync.RWMutex
	m map[string]models.VirtualFile
}{m: make(map[string]models.VirtualFile)}

// FileHandler is the HTTP handler function to serve the virtual directories and files.
func FileHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	virtualFS.RLock()
	file, exists := virtualFS.m[path]
	virtualFS.RUnlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	if file.IsDir {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "<!DOCTYPE html><html><body><ul>")
		virtualFS.RLock()

		// Store eligible items in a map to dedupe them
		children := make(map[string]struct{})
		for subPath, subFile := range virtualFS.m {
			if strings.HasPrefix(subPath, path) && subPath != path {
				remainingPath := strings.TrimPrefix(subPath, path)
				if strings.ContainsRune(remainingPath, '/') {
					// If it contains '/', it's not an immediate child, get the directory part
					dirPart := remainingPath[:strings.Index(remainingPath, "/")]
					children[path+dirPart+"/"] = struct{}{} // Add trailing slash to directory
				} else {
					// It's an immediate child
					if subFile.IsDir {
						children[subPath+"/"] = struct{}{} // Add trailing slash to directory
					} else {
						children[subPath] = struct{}{}
					}
				}
			}
		}
		virtualFS.RUnlock()

		// Loop over the deduped eligible items and display them
		for childPath := range children {
			displayPath := strings.TrimPrefix(childPath, path)
			fmt.Fprintf(w, "<li><a href=\"%s\">%s</a></li>", childPath, displayPath)
		}

		fmt.Fprint(w, "</ul></body></html>")
	} else {
		// Check if the file is already cached locally
		cachePath := filepath.Join(cacheDir, path)
		if _, err := os.Stat(cachePath); os.IsNotExist(err) {
			// If not cached, download and cache the file
			//err := downloadAndStreamFile(w, file.DirectLink, cachePath)
			err := downloadAndStreamFile(w, r, file.DirectLink, cachePath)
			if err != nil {
				log.Printf("Error downloading and streaming file: %v", err) // Log the detailed error message
				http.Error(w, "Failed to download file", http.StatusInternalServerError)
				return
			}
		}

		// Serve the cached file
		http.ServeFile(w, r, cachePath)
	}
}

// AddToFileSystem adds a new file or directory to the virtual file system.
func AddToFileSystem(path string, file models.VirtualFile) {
	virtualFS.Lock()
	defer virtualFS.Unlock()

	// Check if the path already exists in the virtualFS map
	if _, exists := virtualFS.m[path]; !exists {
		virtualFS.m[path] = file
		// log the addition of the file or directory
		//log.Printf("Added %s to virtual file system\n", path)
	}
}

func downloadAndStreamFile(w http.ResponseWriter, r *http.Request, url string, cachePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Set the Content-Type header
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	// Check if a range is requested
	rangeReq := r.Header.Get("Range")
	if rangeReq != "" {
		start, end := parseRange(rangeReq, resp.ContentLength)
		w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, resp.ContentLength))
		w.WriteHeader(http.StatusPartialContent)

		// Discard bytes until the start of the requested range
		if _, err := io.CopyN(io.Discard, resp.Body, start); err != nil {
			return err
		}

		// Limit the number of bytes read from resp.Body to the requested range
		limitedReader := io.LimitReader(resp.Body, end-start+1)
		io.Copy(w, limitedReader)
	} else {
		w.WriteHeader(http.StatusOK)
		io.Copy(w, resp.Body)
	}

	// Create the necessary directories before creating the cache file
	dir := filepath.Dir(cachePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	// Save the file to cache
	cacheFile, err := os.Create(cachePath)
	if err != nil {
		return err
	}
	defer cacheFile.Close()
	_, err = io.Copy(cacheFile, resp.Body)
	return err
}

func parseRange(rangeHeader string, fileSize int64) (start int64, end int64) {
	// Parse the Range header value, e.g., "bytes=1000-2000"
	// This is a simplified example; actual implementation may be more complex
	parts := strings.Split(rangeHeader, "=")
	if len(parts) != 2 || parts[0] != "bytes" {
		return 0, fileSize - 1
	}
	rangeParts := strings.Split(parts[1], "-")
	start, _ = strconv.ParseInt(rangeParts[0], 10, 64)
	if len(rangeParts) > 1 && rangeParts[1] != "" {
		end, _ = strconv.ParseInt(rangeParts[1], 10, 64)
	} else {
		end = fileSize - 1
	}
	return start, end
}

// func downloadAndStreamFile(w http.ResponseWriter, url string, cachePath string) error {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Create the necessary directories before creating the cache file
// 	dir := filepath.Dir(cachePath)
// 	if err := os.MkdirAll(dir, 0755); err != nil {
// 		return fmt.Errorf("failed to create directories: %w", err)
// 	}

// 	cacheFile, err := os.Create(cachePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer func() {
// 		cacheFile.Close()
// 		if err != nil {
// 			// If an error occurred, delete the partially downloaded file
// 			os.Remove(cachePath)
// 		}
// 	}()

// 	// Get the expected content length from the response headers
// 	contentLength, _ := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)

// 	multiWriter := io.MultiWriter(w, cacheFile)

// 	// Keep track of the number of bytes written
// 	written, err := io.Copy(multiWriter, resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	// Check if the number of bytes written is less than the expected content length
// 	if contentLength > 0 && written < contentLength {
// 		return fmt.Errorf("download interrupted: only %d out of %d bytes were written", written, contentLength)
// 	}

// 	return nil
// }
