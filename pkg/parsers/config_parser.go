package parsers

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Conf is internal representation of config.
type Conf struct {
	Description  string             `yaml:"description"`
	CoinsSymbols map[string]float64 `yaml:"coins"`
}

// GetConfiguration retrives user configuration.
func GetConfiguration() Conf {
	absPath, _ := filepath.Abs("../paddy/.cointicker_config.yaml")
	yamlFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Printf("cannot get the yml file: #%v ", err)
	}

	var config Conf

	yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Printf("cannot parse the yml file: #%v ", err)
	}

	return config
}
