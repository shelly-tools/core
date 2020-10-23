package config

import (
	"log"

	"gopkg.in/yaml.v2"
)

var (
	defaults = []byte(`
databasePath: shellyDB.db
imageStorePath: uploads/
ui:
  listenAdress: 0.0.0.0
  listenPort: 8080
debugging:
  router:
    mode: PROD
    appEngine: false
  logging:
    logLevel: DEBUG
`,
	)
)

// Config is the Config struct which contains all fields available in the Config Yaml in the YAML-Format
type Config struct {
	DatabasePath   string `yaml:"databasePath"`
	ImageStorePath string `yaml:"imageStorePath"`
	UI             struct {
		ListenAdress string `yaml:"listenAdress"`
		ListenPort   int    `yaml:"listenPort"`
	} `yaml:"ui"`
	Debugging struct {
		Router struct {
			Mode      string `yaml:"mode"`
			AppEngine bool   `yaml:"appEngine"`
		} `yaml:"router"`
		Logging struct {
			LogLevel string `yaml:"logLevel"`
		} `yaml:"logging"`
	} `yaml:"debugging"`
}

// New creates a new config with all defaults defined
func New(data []byte) (*Config, error) {
	var config Config

	err := yaml.Unmarshal(defaults, &config)

	if err != nil {
		log.Println("Error while load defaults")
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)

	return &config, err
}
