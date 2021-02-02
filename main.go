package main

import (
	"fishingServer/web"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func readConfig(conFile string){
	yamlFile, err := ioutil.ReadFile(conFile)
	if err != nil {
		fmt.Println("ioutil.ReadFile error: ", err)
	}

	var config web.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("yaml.Unmarshal error: ", err)
	}
	web.Configuration = config
}

func main() {
	pathToConfig := os.Args[1]
	readConfig(pathToConfig)
	web.WebServer()
}
