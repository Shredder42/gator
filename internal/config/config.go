package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	file, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	dat, err := os.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(dat, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil

}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", homeDir, configFileName), nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName

	err := write(*cfg)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
