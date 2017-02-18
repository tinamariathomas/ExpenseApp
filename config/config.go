package config

import (
	"io/ioutil"
	"encoding/json"
)

type AppConfig struct{
	Port string `json:"port"`
}

func ReadConfig(configFile string) (*AppConfig, error){
	configFileContents, err := ioutil.ReadFile(configFile)
	if err != nil{
		return nil, err
	}
	appConfig := &AppConfig{}

	err = json.Unmarshal([]byte(configFileContents),appConfig)
	if err != nil{
		return nil, err
	}

	return appConfig, nil
}
