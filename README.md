# About
Xtream Codes is a platform that provides software solutions for IPTV (Internet Protocol Television) streaming services. Legitimate networks and providers use it to manage the delivery of media.

Xtream2strm is a utility tool designed to fetch stream data from Xtream Codes API and generate .strm files compatible with Jellyfin. The tool allows users to specify categories to ignore and include specific movies or series, ensuring a customized and user-friendly experience.

:warning: This project is not to be used for piracy or any other illegal activities. Users of this project are responsible for verifying the source of the media, or use personal instances of Xtream Codes API only. 

## Features

- Fetches VOD and Series stream data from Xtream Codes API.
- Generates Jellyfin compatible .strm files.
- Allows ignoring specific categories through configuration.
- Implements an explicit include list for movies/series (for a more tailored media center).
- Ensures sanitized and valid file names for the generated .strm files.
- Supports fuzzy search to find movies or series by name.

## Prerequisites

- Go 1.20 or higher

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/kennyparsons/xtream2strm.git
   ```
2. ```sh
   cd xtream2strm
   ```
3. ```sh 
   go build .
    ```

## Usage

1. Edit the `config.yaml` file in the project root directory with your Xtream Codes API details and desired output directory.
   ```yaml
   api_endpoint: "http://your-api-endpoint:port" #API endpoint with no trailing slash
   username: "your-username"
   password: "your-password"
   output_dir: "path/to/output/directory"
   ignore_categories:
   - "category-id-1"
   - "category-id-2"
   movie_include:
   - "movie-id-1"
   - "movie-id-2"
   series_include:
   - "series-id-1"
   - "series-id-2"
   ```
2. Run the tool.
   ```sh
   ./xtream2strm --config path/to/config.yaml
   ```
3. To search for a specific movie or series:
   ```sh
   ./xtream2strm --config path/to/config.yaml --search "Movie or Series Name" 
   ```