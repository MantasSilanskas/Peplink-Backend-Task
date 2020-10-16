package readingData

import (
	"encoding/json"
	"github.com/MantasSilanskas/Peplink-Backend-Task/pkg/rawDataStruct"
	"io/ioutil"
	"os"
)

func ReadDataFile(fileName string) (rawDataStruct.CryptoCurrencyData, error) {

	var data rawDataStruct.CryptoCurrencyDataSlice

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return rawDataStruct.CryptoCurrencyData{}, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return rawDataStruct.CryptoCurrencyData{}, err
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return rawDataStruct.CryptoCurrencyData{}, err
	}

	return data[0], err
}
