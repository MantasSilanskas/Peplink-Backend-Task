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

func Parse(resultMap map[int]bool, rulesPrices map[int]float64) (*map[int]bool, *map[int]float64, error) {


	ruleSet, err := LoadRuleSets("rulesFile.json") // Loads all rule sets from rulesFile.json
	if err != nil {
		return &resultMap, &rulesPrices, err
	}

	dataMap := make(map[string]CryptoCurrencyData)

	for i, v := range ruleSet.Rules {

		fileUrl := baseUrl + v.CryptoID
		if mapData, ok := dataMap[v.CryptoID]; !ok {
			err := DownloadFile(beginningDataFileName+v.CryptoID+dataFileExtension, fileUrl)
			if err != nil {
				return &resultMap, &rulesPrices, err
			}
			mapData, err = ReadDataFile(beginningDataFileName + v.CryptoID + dataFileExtension)
			if err != nil {
				return &resultMap, &rulesPrices, err
			}
			dataMap[v.CryptoID] = mapData
		}
		if rulesPrices[i] != v.Price {
			resultMap[i] = false
		}
		price, err := strconv.ParseFloat(dataMap[v.CryptoID].PriceUsd, 64)
		if err != nil {
			return &resultMap, &rulesPrices, err
		}
		if ok := resultMap[i]; !ok {
			if price > v.Price && v.Rule == "gt" && resultMap[i] == false {
				fmt.Println("Cryptocurrency", " id:", dataMap[v.CryptoID].ID, dataMap[v.CryptoID].Name, "price is greater than", v.Price)
				resultMap[i] = true
			}
			if price < v.Price && v.Rule == "lt" && resultMap[i] == false {
				fmt.Println("Cryptocurrency", " id:", dataMap[v.CryptoID].ID, dataMap[v.CryptoID].Name, "price is lower than", v.Price)
				resultMap[i] = true
			}
		}
		rulesPrices[i] = v.Price
		i++
	}

	return &resultMap, &rulesPrices, nil
}
