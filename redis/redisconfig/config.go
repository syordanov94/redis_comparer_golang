package redisconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type RedisConfig struct {
	Host string `json:"redis_host"`
	Port int    `json:"redis_port"`
	Pass string `json:"redis_password"`
}

func ConfigFromFile(fileName string) (RedisConfig, error) {

	// --- (1) ----
	// Open file containing the config
	configRaw, err := os.Open(fileName)
	if err != nil {
		return RedisConfig{}, err
	}

	// --- (2) ----
	// Parse raw data to JSON struct
	byteValue, err := ioutil.ReadAll(configRaw)
	if err != nil {
		return RedisConfig{}, err
	}

	var config RedisConfig
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return RedisConfig{}, err
	}
	return config, nil
}
