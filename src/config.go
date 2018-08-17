package main

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type TestGenieConfig struct {
	Host struct {
		URL     string `yaml:"url"`
		Live    bool   `yaml:"live"`
		Refresh bool   `yaml:"refresh"`
	}
}

func getConfig() TestGenieConfig {
	config := TestGenieConfig{}

	filename, _ := filepath.Abs("./config.yml")

	yamlFile, err1 := ioutil.ReadFile(filename)
	errorHandler(err1)

	err2 := yaml.Unmarshal(yamlFile, &config)
	errorHandler(err2)

	return config
}
