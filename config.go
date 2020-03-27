package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var sharedConfig Config

// Config is initialized from asuite.yaml
type Config struct {
	OrganisationID string `yaml:"organisationID"`
	APIKey         string `yaml:"api-key"`
	DirectoryID    string `yaml:"directoryID"`
}

func init() {
	data, err := ioutil.ReadFile("asuite.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(data, &sharedConfig); err != nil {
		log.Fatal(err)
	}
}
