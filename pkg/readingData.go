package pkg

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadDataFile(fileName string) (CryptoCurrencyData, error) {

	var data CryptoCurrencyData

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return data, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return data, err
	}

	return data, err
}
