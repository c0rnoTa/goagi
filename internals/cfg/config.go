package cfg

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Configuration struct {
	App struct {
		Debug bool `yaml:"debug"` // TODO: перевести на logrus
	} `yaml:"app"`
	Database struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

// Load file contents to Configuration object.
func (c *Configuration) Load(configFile string) error {
	// TODO: дефаолтовые значения, viper

	file, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(file, c)
}
