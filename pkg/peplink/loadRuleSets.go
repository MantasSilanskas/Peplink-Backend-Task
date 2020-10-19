package peplink

import (
	"encoding/json"
	"io/ioutil"
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

	// visa funkciaj galiam sutrumpinti pasinaudojus ioutil.ReadFile
	var (
		err  error
		r    ruleSets
		data []byte
	)
	if data, err = ioutil.ReadFile(filename); err != nil {
		return ruleSets{}, err
	}

	if err = json.Unmarshal(data, &r); err != nil {
		return ruleSets{}, err
	}
	return r, err
}
