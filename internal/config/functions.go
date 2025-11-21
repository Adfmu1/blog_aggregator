package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func CheckPrintErr(err error) {
	if err != nil {
		fmt.Println("an error has occured:", err)
	}
}

func Read() Config {
	conf := Config{}
	// get config file path
	confPath, err := os.UserHomeDir()
	CheckPrintErr(err)
	confFilePath := confPath + "/gatorconfig.json"
	// get config file contents
	configData, err := os.ReadFile(confFilePath)
	CheckPrintErr(err)
	err = json.Unmarshal(configData, &conf)
	CheckPrintErr(err)
	return conf
}
