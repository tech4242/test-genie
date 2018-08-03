package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type TestGenieConfig struct {
	Host struct {
		Url  string `yaml:"url"`
		Live bool   `yaml:"live"`
	}
}

func get_config() TestGenieConfig {
	config := TestGenieConfig{}

	filename, _ := filepath.Abs("./config.yml")
	yamlFile, err1 := ioutil.ReadFile(filename)

	if err1 != nil {
		log.Fatalf("error: %v", err1)
	}

	err2 := yaml.Unmarshal(yamlFile, &config)

	if err2 != nil {
		log.Fatalf("error: %v", err2)
	}

	return config
}
