package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// Config is initialized from asuite.yaml
type Config struct {
	OrganisationID string `yaml:"organisationID"`
	APIKey         string `yaml:"api-key"`
	DirectoryID    string `yaml:"directoryID"`
}

func getConfig(c *cli.Context) (cfg Config) {
	configfile := c.GlobalString("config")
	if len(configfile) == 0 {
		if info, err := os.Stat("asuite.yaml"); err == nil && !info.IsDir() {
			configfile = "asuite.yaml"
		}
	}
	if len(configfile) == 0 {
		configfile = filepath.Join(os.Getenv("HOME"), "asuite.yaml")
	}
	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatal(err)
	}
	return
}
