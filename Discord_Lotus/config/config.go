package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

var (
	Token     string
	BotPrefix string
	config    *configStruct
)

func ReadConfig() error {
	fmt.Println()

	file, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println("error detected when reading file")
		return err
	}
	fmt.Println(string(file))
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("error when parsing configurations %v", err)
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil

}
