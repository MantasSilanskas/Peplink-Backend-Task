package results

import (
	"fmt"
	"github.com/MantasSilanskas/Peplink-Backend-Task/pkg/download"
	"github.com/MantasSilanskas/Peplink-Backend-Task/pkg/loadRuleSets"
	"github.com/MantasSilanskas/Peplink-Backend-Task/pkg/rawDataStruct"
	"github.com/MantasSilanskas/Peplink-Backend-Task/pkg/readingData"
	"strconv"
)

const (
	baseUrl               = "https://api.coinlore.com/api/ticker/?id="
	beginningDataFileName = "rawData"
	dataFileExtension     = ".json"
)

func Parse() error {

	ruleSet, err := loadRuleSets.LoadRuleSets("rulesFile.json") // Loads all rule sets from rulesFile.json
	if err != nil {
		fmt.Println(err)
	}

	dataMap := make(map[string]rawDataStruct.CryptoCurrencyData)

	for i, v := range ruleSet.Rules {
		fileUrl := baseUrl + v.CryptoID
		if mapData, ok := dataMap[v.CryptoID]; !ok {
			dataFile, err := download.DownloadFile(beginningDataFileName+v.CryptoID+dataFileExtension, fileUrl)
			if err != nil {
				return err
			}
			mapData, err = readingData.ReadDataFile(dataFile)
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
