package models

type Config struct {
    APIEndpoint       string   `yaml:"api_endpoint"`
    Username          string   `yaml:"username"`
    Password          string   `yaml:"password"`
    OutputDir         string   `yaml:"output_dir"`
    IgnoreCategories  []string `yaml:"ignore_categories"`
    MovieInclude      []string `yaml:"movie_include"`      // Whitelist for movies
    SeriesInclude     []string `yaml:"series_include"`     // Whitelist for series
}