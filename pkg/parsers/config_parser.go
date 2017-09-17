package parsers

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Conf is internal representation of config.
type Conf struct {
	Coins []string `yaml:"coins"`
}

// GetConfiguration retrives user configuration.
func (c *Conf) GetConfiguration() *Conf {
	absPath, _ := filepath.Abs("../paddy/.cointicker_config.yaml")
	yamlFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Printf("cannot get the yml file: #%v ", err)
	}

	yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("cannot parse the yml file: #%v ", err)
	}

	return c
}
