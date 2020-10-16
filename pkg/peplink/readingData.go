package peplink

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadDataFile(fileName string) (CryptoCurrencyData, error) {

	var data CryptoCurrencyDataSlice

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return CryptoCurrencyData{}, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return CryptoCurrencyData{}, err
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return CryptoCurrencyData{}, err
	}

	return data[0], err
}
