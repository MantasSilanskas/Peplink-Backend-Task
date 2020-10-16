package peplink

import (
	"fmt"
	"strconv"
)

const (
	baseUrl               = "https://api.coinlore.com/api/ticker/?id="
	beginningDataFileName = "rawData"
	dataFileExtension     = ".json"
)

func Parse() error {

	ruleSet, err := LoadRuleSets("rulesFile.json") // Loads all rule sets from rulesFile.json
	if err != nil {
		return err
	}

	dataMap := make(map[string]CryptoCurrencyData)

	for i, v := range ruleSet.Rules {
		fileUrl := baseUrl + v.CryptoID
		if mapData, ok := dataMap[v.CryptoID]; !ok {
			err := DownloadFile(beginningDataFileName+v.CryptoID+dataFileExtension, fileUrl)
			if err != nil {
				return err
			}
			mapData, err = ReadDataFile(beginningDataFileName + v.CryptoID + dataFileExtension)
			if err != nil {
				return err
			}
			dataMap[v.CryptoID] = mapData
		}
		price, err := strconv.ParseFloat(dataMap[v.CryptoID].PriceUsd, 64)
		if err != nil {
			return err
		}

		if price > v.Price && v.Rule == "gt" {
			fmt.Println("Cryptocurrency", dataMap[v.CryptoID].ID, dataMap[v.CryptoID].Name, "price is greater than", v.Price)
		}
		if price < v.Price && v.Rule == "lt" {
			fmt.Println("Cryptocurrency", dataMap[v.CryptoID].ID, dataMap[v.CryptoID].Name, "price is lower than", v.Price)
		}
		i++
	}

	return nil
}
