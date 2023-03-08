package app

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	SSHFile     string `json:"sshfile"`
	ResultsPath string `json:"resultspath"`
	HostsPath   string `json:"hostspath"`
	CVEsPath    string `json:"cvespath"`
}

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}
