package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func init() {
	loadConf()
}

type cweb struct {
	Nsqdtcpaddr string `yaml:"nsqdtcpaddr"`
	// Imgserveraddr string `yaml:"imgserveraddr"`
}

type Services struct {
	Web cweb `yaml:"web"`
}
type Conf struct {
	Version  string   `yaml:"version"`
	Services Services `yaml:"services"`
}

var Config Conf

func loadConf() {
	yamlFile, err := ioutil.ReadFile("./conf.yml")
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("yamlFile.Unmarshal: %v", err)
	}
	//
	//f, err := yaml.Marshal(Config)
	//ioutil.WriteFile("./t3.yml", f, 0666)
}
