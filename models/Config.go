package models

type Config struct {
	APIEndpoint      string   `yaml:"api_endpoint"`
	Username         string   `yaml:"username"`
	Password         string   `yaml:"password"`
	OutputDir        string   `yaml:"output_dir"`
	IgnoreCategories []string `yaml:"ignore_categories"`
	MovieInclude     []int    `yaml:"movie_include"`  // Whitelist for movies
	SeriesInclude    []int    `yaml:"series_include"` // Whitelist for series
	NotFound         struct {
		Movies []string `yaml:"movies"`
		TV     []string `yaml:"tv"`
	} `yaml:"notfound"`
}
