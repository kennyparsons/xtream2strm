package models

// Config represents the structure of the configuration file.
type Config struct {
	APIEndpoint      string   `yaml:"api_endpoint"`
	Username         string   `yaml:"username"`
	Password         string   `yaml:"password"`
	OutputDir        string   `yaml:"output_dir"`
	IgnoreCategories []string `yaml:"ignore_categories"`
	// Add any other configuration fields as needed.
}
