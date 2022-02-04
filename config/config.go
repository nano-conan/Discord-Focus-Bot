package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string
	UserID	string
	ServerID	string
	Start	int
	End		int

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	UserID	string	`json:"UserID"`
	ServerID	string	`json:"ServerID"`
	Start		int	`json:"Start"`
	End			int	`json:"End"`
}

func ReadConfig() error {
	fmt.Println("Reading config file")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	UserID = config.UserID
	ServerID = config.ServerID
	Start = config.Start
	End = config.End

	return nil

}
