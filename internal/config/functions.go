package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// ============ CONSTANTS ============
const configFileName = "/gatorconfig.json"

// ============ HELPER FUNCTIONS ============
func CheckPrintErr(err error) {
	if err != nil {
		fmt.Println("an error has occured:", err)
	}
}

// ============ MAIN FUNCTIONS ============
// READ THE CONFIG FILE
func Read() Config {
	// get config file path
	confPath, err := os.UserHomeDir()
	CheckPrintErr(err)
	confFilePath := confPath + configFileName
	// get config file contents
	configData, err := os.ReadFile(confFilePath)
	CheckPrintErr(err)
	// unmarshal the data
	conf := Config{}
	err = json.Unmarshal(configData, &conf)
	CheckPrintErr(err)
	return conf
}

// WRITE NEW USER TO CONFIG FILE
func (conf Config) SetUser(userName string) {
	// set new username in struct
	conf.CurrentUser = userName
	// get config file path
	confPath, err := os.UserHomeDir()
	CheckPrintErr(err)
	confFilePath := confPath + configFileName
	// marshal the data
	data, err := json.Marshal(conf)
	CheckPrintErr(err)
	err = os.WriteFile(confFilePath, data, 0666)
	CheckPrintErr(err)
}
