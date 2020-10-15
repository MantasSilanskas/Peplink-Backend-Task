package pkg

import (
	"encoding/json"
	"os"
)

type ruleSets struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Rules       []struct {
		CryptoID string  `json:"crypto_id"`
		Price    float64 `json:"price"`
		Rule     string  `json:"rule"`
	} `json:"rules"`
}

func LoadRuleSets(filename string) (ruleSets, error) {

	var ruleSet ruleSets
	ruleSetFile, err := os.Open(filename)

	defer ruleSetFile.Close()
	if err != nil {
		return ruleSet, err
	}

	jsonParser := json.NewDecoder(ruleSetFile)
	err = jsonParser.Decode(&ruleSet)
	if err != nil {
		return ruleSet, err
	}

	return ruleSet, err
}
