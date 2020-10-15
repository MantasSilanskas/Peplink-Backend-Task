package pkg

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
		fmt.Println(err)
	}

	m := make(map[string]CryptoCurrencyData)

	for _, v := range ruleSet.Rules {
		fileUrl := baseUrl + v.CryptoID
		if mapData, ok := m[v.CryptoID]; !ok {
			dataFile, err := DownloadFile(beginningDataFileName+v.CryptoID+dataFileExtension, fileUrl)
			if err != nil {
				return err
			}
			mapData, err = ReadDataFile(dataFile)
			if err != nil {
				return err
			}
			m[v.CryptoID] = mapData
		}
		price, err := strconv.ParseFloat(m[v.CryptoID].PriceUsd, 64)
		if err != nil {
			return err
		}

		if price > v.Price && v.Rule == "gt" {
			fmt.Println("Cryptocurrency", m[v.CryptoID].ID, m[v.CryptoID].Name, "price is greater than", v.Price)
		}
		if price < v.Price && v.Rule == "lt" {
			fmt.Println("Cryptocurrency", m[v.CryptoID].ID, m[v.CryptoID].Name, "price is lower than", v.Price)
		}
	}

	return nil
}
