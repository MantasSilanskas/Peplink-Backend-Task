package peplink

import (
	"fmt"
	"strconv"
	"time"
)

const (
	baseUrl               = "https://api.coinlore.com/api/ticker/?id="
	beginningDataFileName = "rawData"
	dataFileExtension     = ".json"
)

func Parse(isPrinted [...]bool) ([...]bool, error) {

	printed := isPrinted
	ruleSet, err := LoadRuleSets("rulesFile.json") // Loads all rule sets from rulesFile.json
	if err != nil {
		return printed, err
	}

	dataMap := make(map[string]CryptoCurrencyData)

	for i, v := range ruleSet.Rules {
		fileUrl := baseUrl + v.CryptoID
		if mapData, ok := dataMap[v.CryptoID]; !ok {
			err := DownloadFile(beginningDataFileName+v.CryptoID+dataFileExtension, fileUrl)
			if err != nil {
				return printed, err
			}
			mapData, err = ReadDataFile(beginningDataFileName + v.CryptoID + dataFileExtension)
			if err != nil {
				return printed, err
			}
			dataMap[v.CryptoID] = mapData
		}
		price, err := strconv.ParseFloat(dataMap[v.CryptoID].PriceUsd, 64)
		if err != nil {
			return printed, err
		}
		if price > v.Price && v.Rule == "gt" && printed[i] == false {
			fmt.Println(time.Stamp)
			fmt.Println("Cryptocurrency", " id:", dataMap[v.CryptoID].ID, dataMap[v.CryptoID].Name, "price is greater than", v.Price)
			printed[i] = true
		}
		if price < v.Price && v.Rule == "lt" && printed[i] == false {
			fmt.Println(time.Stamp)
			fmt.Println("Cryptocurrency", " id:", dataMap[v.CryptoID].ID, dataMap[v.CryptoID].Name, "price is lower than", v.Price)
			printed[i] = true
		}
		i++
	}

	return printed, nil
}
