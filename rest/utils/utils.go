package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/go-feature-flag/rest/model"
)

func GetFeatureFlagKeys(m map[string]model.FeatureFlagConfig) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func GetFeatureFlagConfig(configFilePath string) (map[string]model.FeatureFlagConfig, error) {
	var result map[string]model.FeatureFlagConfig

	// Open our jsonFile
	jsonFile, err := os.Open(configFilePath)
	if err != nil {
		return result, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal([]byte(byteValue), &result); err != nil {
		return result, err
	}

	return result, nil
}

func OverwriteFeatureFlagConfig(configFilePath string, config map[string]model.FeatureFlagConfig) error {
	newFile, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return errors.New("failed to open new file config")
	}
	defer newFile.Close()

	encoder := json.NewEncoder(newFile)
	encoder.Encode(config)

	return nil
}
